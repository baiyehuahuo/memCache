package main

import (
	"memCache/cache"
	"time"
)

func main() {
	c := cache.NewMemCache()
	expiration := time.Second
	c.SetMaxMemory("2MB")
	c.SetMaxMemory("2kb")
	c.SetMaxMemory("20")
	c.SetMaxMemory("2Z")
	c.Set("int", 1, expiration)
	c.Set("bool", false, expiration)
	c.Set("data", map[string]any{"a": "1"}, expiration)
	c.Get("int")
	c.Del("int")
	c.Flush()
	c.Keys()
}
