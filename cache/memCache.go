package cache

import (
	"fmt"
	"memCache/define"
	"memCache/util"
	"sync"
	"time"
)

var _ Cache = (*memCache)(nil)

type memCache struct {
	// 线程锁
	mutex sync.RWMutex

	// 最大内存
	maxMemorySize int64

	// 最大内存的字符串表示
	maxMemorySizeStr string

	// 当前已使用内存
	currentMemorySize int64

	// 缓存键值对的映射表
	values map[string]*memCacheValue
}

type memCacheValue struct {
	// 实际值
	value any
	// 过期时间，绝对时间
	expireTime time.Time
	// value 大小
	size int64
}

func NewMemCache() Cache {
	return &memCache{
		maxMemorySize:     define.DefaultMemSize,
		maxMemorySizeStr:  define.DefaultMemSizeStr,
		currentMemorySize: 0,
		values:            make(map[string]*memCacheValue),
	}
}

func (mc *memCache) SetMaxMemory(size string) bool {
	fmt.Println("setMaxMemory", size)
	mc.maxMemorySize, mc.maxMemorySizeStr = util.ParseSize(size)
	return false
}

func (mc *memCache) set(key string, val *memCacheValue) {
	mc.values[key] = val
	mc.currentMemorySize += val.size
}

func (mc *memCache) del(key string) {
	val, exists := mc.values[key]
	if !exists {
		return
	}
	mc.currentMemorySize -= val.size
	delete(mc.values, key)
}

func (mc *memCache) Set(key string, val any, expiration time.Duration) bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	fmt.Println("set", key, val, expiration)

	var previousSize int64
	if v, ok := mc.values[key]; ok {
		previousSize = v.size
	}

	size := util.GetValueSize(val)
	if mc.currentMemorySize+size-previousSize > mc.maxMemorySize {
		return false
	}

	mcv := &memCacheValue{
		value:      val,
		expireTime: time.Now().Add(expiration),
		size:       size,
	}
	mc.values[key] = mcv
	mc.currentMemorySize += size - previousSize
	return true
}

func (mc *memCache) Get(key string) (val any, exists bool) {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	fmt.Println("get", key)

	mcv, exists := mc.values[key]
	if !exists {
		return nil, false
	}

	if mcv.expireTime.Before(time.Now()) {
		mc.del(key)
		return nil, false
	}

	return mcv.value, true
}

func (mc *memCache) Del(key string) {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	fmt.Println("del", key)

	mc.del(key)
}

func (mc *memCache) Exists(key string) bool {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	fmt.Println("exists", key)

	mcv, exists := mc.values[key]
	if !exists {
		return false
	}

	if mcv.expireTime.Before(time.Now()) {
		mc.del(key)
		return false
	}

	return true
}

func (mc *memCache) Flush() {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	fmt.Println("flush")

	mc.currentMemorySize = 0
	mc.values = make(map[string]*memCacheValue)
}

func (mc *memCache) Keys() int64 {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	fmt.Println("keys")

	var counter int64
	for key, val := range mc.values {
		if val.expireTime.Before(time.Now()) {
			mc.del(key)
			continue
		}
		counter++
	}

	return counter
}
