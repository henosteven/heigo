package lib

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/henosteven/heigo/config"
	"net"
	"time"
)

var pool *redis.Pool

func InitRedis(config config.RedisConfig) {
	pool = &redis.Pool{
		MaxIdle:     config.MaxIdle,
		IdleTimeout: config.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", net.JoinHostPort(config.Host, config.Port))
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("ping")
			if err != nil {
				return fmt.Errorf("")
			}
			return err
		},
	}
}

func Set(key, val string) error {
	_, err := pool.Get().Do("Set", key, val)
	return err
}

func Get(key string) (string, error) {
	val, err := redis.String(pool.Get().Do("Get", key))
	return val, err
}
