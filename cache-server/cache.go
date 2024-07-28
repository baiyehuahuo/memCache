package cache_server

import (
	"memCache/cache"
	"time"
)

type CacheServer struct {
	memCache cache.Cache
}

func NewCacheServer() *CacheServer {
	mc := cache.NewMemCache()
	return &CacheServer{
		memCache: mc,
	}
}

func (cs *CacheServer) SetMaxMemory(size string) bool {
	return cs.memCache.SetMaxMemory(size)
}

func (cs *CacheServer) Set(key string, val any, expires ...time.Duration) bool {
	if len(expires) > 0 {
		return cs.memCache.Set(key, val, expires[0])
	}
	return cs.memCache.Set(key, val, 0)
}

func (cs *CacheServer) Get(key string) (val any, exists bool) {
	return cs.memCache.Get(key)
}

func (cs *CacheServer) Del(key string) {
	cs.memCache.Del(key)
}

func (cs *CacheServer) Exists(key string) bool {
	return cs.memCache.Exists(key)
}

func (cs *CacheServer) Flush() {
	cs.memCache.Flush()
}

func (cs *CacheServer) Keys() int64 {
	return cs.memCache.Keys()
}
