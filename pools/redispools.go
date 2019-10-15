package pools


import (
	"os"
	"redigo/redis"
	"time"
)


var RD *redis.Pool

/**
redis 连接池
 */
func init(){
	addr := os.Getenv("REDIS_ADDR")
	r := &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp",addr)
		},
		// 最大空闲
		MaxIdle:         50,
		//当为零时，池中的连接数没有限制。
		MaxActive:       0,
		// 超时时间
		IdleTimeout:     240 * time.Second,
		// 当超过最大连接数时直接返回错误 默认false
		Wait:            false,
		// 最大等待时间 默认为0 无限制
		MaxConnLifetime: 0,
	}
	RD = r
}
