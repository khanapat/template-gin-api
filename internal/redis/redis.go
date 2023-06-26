package redis

import (
	"template-gin-api/config"

	"github.com/gomodule/redigo/redis"
)

func NewRedisConn(cfg *config.Config) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     cfg.Redis.MaxIdle,
		IdleTimeout: cfg.Redis.Timeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cfg.Redis.Host)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", cfg.Redis.Password); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}
