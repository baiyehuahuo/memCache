package cache

import (
	"time"
)

type memCacheValue struct {
	// 实际值
	value any
	// 过期时间，绝对时间
	expireTime time.Time
	// value 大小
	size int64
}

func (mcv *memCacheValue) isExpired() bool {
	// 非空且超时
	return !mcv.expireTime.IsZero() && time.Now().After(mcv.expireTime)
}
