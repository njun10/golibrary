package redis

import (
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/frame/g"
)


var (
	CachePool = make(map[string]*gredis.Redis)
)

//初始化redis 连接池
func Init(name string) *gredis.Redis {
	if c, ok := CachePool[name]; ok{
		return c
	}
	CachePool[name] = g.Redis(name)
	defer CachePool[name].Close()
	return CachePool[name]
}
