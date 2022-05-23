package redis

import (
	string2 "go-sv/tools/redis/string"
	"time"
)

type DBFunc func() string

type Cache struct {
	Operation *string2.StringOperation
	Expire    time.Duration
	DB        DBFunc
}

// 创建缓存
func NewCache(operation *string2.StringOperation, expr time.Duration) *Cache {
	return &Cache{
		Operation: operation,
		Expire:    expr,
	}
}

// 设置缓存
func (c *Cache) SetCache(key string, value interface{}) {
	c.Operation.Set(key, value, string2.WithExpire(c.Expire)).UnWarp()
}

func (c *Cache) GetCache(key string) {
	ret := c.Operation.Get(key).UnWarpOrElse(c.DB)
	c.SetCache(key, ret)
	return
}
