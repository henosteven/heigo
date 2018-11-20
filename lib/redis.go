package lib

import (
	"github.com/garyburd/redigo/redis"
	"github.com/henosteven/heigo/config"
	"net"
	"time"
	"fmt"
)

var pool *redis.Pool

func InitRedis (config config.RedisConfig) {
	pool = &redis.Pool{
		MaxIdle:config.MaxIdle,
		IdleTimeout:config.IdleTimeout,
		Dial: func() (redis.Conn, error){
			c , err := redis.DialURL(net.JoinHostPort(config.Host, config.Password))
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
