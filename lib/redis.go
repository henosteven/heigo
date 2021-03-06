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

			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("ping")
			if err != nil {
				return fmt.Errorf("")
			}
			return err
		},
	}
}

func Set(key, val string) error {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("Set", key, val)
	return err
}

func Get(key string) (string, error) {
	conn := pool.Get()
	defer conn.Close()
	reply, err := conn.Do("Get", key)
	val, err := redis.String(reply, err)
	return val, err
}
