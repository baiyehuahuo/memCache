package cache

import (
	"memCache/define"
	"memCache/util"
	"sync"
	"time"
)

var _ Cache = (*memCache)(nil)

type memCache struct {
	// 线程锁
	mutex sync.Mutex

	// 最大内存
	maxMemorySize int64

	// 最大内存的字符串表示
	maxMemorySizeStr string

	// 当前已使用内存
	currentMemorySize int64

	// 缓存键值对的映射表
	values map[string]*memCacheValue

	// 自动触发清理过期函数的时间间隔
	clearExpiredTimeInterval time.Duration
}

func NewMemCache() Cache {
	mc := &memCache{
		maxMemorySize:            define.DefaultMemSize,
		maxMemorySizeStr:         define.DefaultMemSizeStr,
		currentMemorySize:        0,
		values:                   make(map[string]*memCacheValue),
		clearExpiredTimeInterval: time.Second,
	}

	go mc.autoCleanExpiredItems()

	return mc
}

// 添加键值对 外部上锁后再使用
func (mc *memCache) set(key string, val *memCacheValue) {
	mc.values[key] = val
	mc.currentMemorySize += val.size
}

// 删除键值对 外部上锁后再使用
func (mc *memCache) del(key string) {
	val, exists := mc.values[key]
	if !exists {
		return
	}
	mc.currentMemorySize -= val.size
	delete(mc.values, key)
}

// 清理过期键值对，外部上锁后再使用
func (mc *memCache) cleanExpiredItems() {
	for key, val := range mc.values {
		if val.isExpired() {
			mc.del(key)
		}
	}
}

// 定时启动清理的线程函数
func (mc *memCache) autoCleanExpiredItems() {
	ticker := time.NewTicker(mc.clearExpiredTimeInterval)
	for {
		select {
		case <-ticker.C:
			mc.mutex.Lock()
			mc.cleanExpiredItems()
			mc.mutex.Unlock()
		}
	}
}

func (mc *memCache) SetMaxMemory(size string) bool {
	mc.maxMemorySize, mc.maxMemorySizeStr = util.ParseSize(size)
	return false
}

func (mc *memCache) Set(key string, val any, expire time.Duration) bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	var previousSize int64
	if v, ok := mc.values[key]; ok {
		previousSize = v.size
	}

	size := util.GetValueSize(val)
	if mc.currentMemorySize+size-previousSize > mc.maxMemorySize {
		return false
	}

	mcv := &memCacheValue{
		value: val,
		size:  size,
	}
	if expire != 0 {
		mcv.expireTime = time.Now().Add(expire)
	}

	mc.values[key] = mcv
	mc.currentMemorySize += size - previousSize
	return true
}

func (mc *memCache) Get(key string) (val any, exists bool) {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	mcv, exists := mc.values[key]
	if !exists {
		return nil, false
	}

	if mcv.isExpired() {
		mc.del(key)
		return nil, false
	}

	return mcv.value, true
}

func (mc *memCache) Del(key string) {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	mc.del(key)
}

func (mc *memCache) Exists(key string) bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	mcv, exists := mc.values[key]
	if !exists {
		return false
	}

	if mcv.isExpired() {
		mc.del(key)
		return false
	}

	return true
}

func (mc *memCache) Flush() {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	mc.currentMemorySize = 0
	mc.values = make(map[string]*memCacheValue)
}

func (mc *memCache) Keys() int64 {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	mc.cleanExpiredItems()

	return int64(len(mc.values))
}
