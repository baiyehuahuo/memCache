package cache

import (
	"fmt"
	"time"
)

var _ Cache = (*memCache)(nil)

type memCache struct {
}

func NewMemCache() Cache {
	return &memCache{}
}

func (m *memCache) SetMaxMemory(size string) bool {
	fmt.Println("setMaxMemory")
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
