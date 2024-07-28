package cache

import (
	"fmt"
	"memCache/util"
	"time"
)

var _ Cache = (*memCache)(nil)

type memCache struct {
	// 最大内存
	maxMemorySize int64

	// 最大内存的字符串表示
	maxMemorySizeStr string

	// 当前已使用内存
	currentMemorySize int64
}

func NewMemCache() Cache {
	return &memCache{}
}

func (m *memCache) SetMaxMemory(size string) bool {
	fmt.Println("setMaxMemory", size)
	m.maxMemorySize, m.maxMemorySizeStr = util.ParseSize(size)
	return false
}

func (m *memCache) Set(key string, val any, expiration time.Duration) bool {
	fmt.Println("set", key, val, expiration)
	return false
}

func (m *memCache) Get(key string) (val any, exists bool) {
	fmt.Println("get", key)
	return nil, false
}

func (m *memCache) Del(key string) bool {
	fmt.Println("del", key)
	return false
}

func (m *memCache) Exists(key string) bool {
	fmt.Println("exists", key)
	return false
}

func (m *memCache) Flush() bool {
	fmt.Println("flush")
	return false
}

func (m *memCache) Keys() int64 {
	fmt.Println("keys")
	return 0
}
