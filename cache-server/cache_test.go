package cache_server

import (
	"strconv"
	"testing"
	"time"
)

func TestSetGet(t *testing.T) {
	cs := NewCacheServer()
	key, val := "int", 1
	cs.Set(key, val)
	v, exists := cs.Get(key)
	if !exists {
		t.Error("key not found")
	}
	if v.(int) != val {
		t.Error("value mismatch")
	}
}

func TestSetGetExpire(t *testing.T) {
	cs := NewCacheServer()
	key, val, expireTime := "int", 1, time.Second
	cs.Set(key, val, expireTime)

	time.Sleep(expireTime)

	_, exists := cs.Get(key)
	if exists {
		t.Error("expired key found")
	}
}

func TestDel(t *testing.T) {
	cs := NewCacheServer()
	key, val := "int", 1
	cs.Set(key, val)
	cs.Del(key)

	_, exists := cs.Get(key)
	if exists {
		t.Error("expired key found")
	}
}

func TestKeys(t *testing.T) {
	cs := NewCacheServer()
	var n = 10
	for i := range n {
		cs.Set(strconv.Itoa(i), i)
	}
	keys := cs.Keys()
	if int(keys) != n {
		t.Error("keys mismatch")
	}
}

func TestFlush(t *testing.T) {
	cs := NewCacheServer()
	var n = 10
	for i := range n {
		cs.Set(strconv.Itoa(i), i)
	}
	cs.Flush()
	keys := cs.Keys()
	if keys != 0 {
		t.Error("flush failed")
	}
}
