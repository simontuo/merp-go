package redis

import (
	"fmt"
	"time"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/simontuo/merp-go/config"
)

var RedisPool *redigo.Pool

func Init() *redigo.Pool {
	var server string
	server = fmt.Sprintf("%s:%s", config.G_Redis.Host, config.G_Redis.Port)

	return PoolInitRedis(server, config.G_Redis.Password)
}

func GetRedisPool() *redigo.Pool {
	return Init()
}

// redis pool
func PoolInitRedis(server string, password string) *redigo.Pool {
	return &redigo.Pool{
		MaxIdle:     2, //空闲数
		IdleTimeout: 240 * time.Second,
		MaxActive:   3, //最大数
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
