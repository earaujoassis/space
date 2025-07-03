package memory

import (
	"fmt"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/gomodule/redigo/redis"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/logs"
)

func NewRedisProviderPool(cfg *config.Config) *redis.Pool {
	switch cfg.Environment {
	case config.Production, config.Development, config.Integration:
		return &redis.Pool{
			MaxIdle:     10,
			MaxActive:   100,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.DialURL(getRedisURL(cfg))
				if err != nil {
					logs.Propagate(logs.LevelError, err.Error())
					return nil, err
				}

				return c, nil
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				if err != nil {
					logs.Propagate(logs.LevelError, err.Error())
				}
				return err
			},
		}
	case config.Test:
		provider, err := miniredis.Run()
		if err != nil {
			logs.Propagate(logs.LevelPanic, err.Error())
			return nil
		}
		return &redis.Pool{
			MaxIdle: 3,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", provider.Addr())
			},
		}
	default:
		logs.Propagate(logs.LevelPanic, "gateway misconfigured")
		return nil
	}
}

func getRedisURL(cfg *config.Config) string {
	switch cfg.Environment {
	case config.Production:
		return fmt.Sprintf("redis://:%s@%s:%d/%d",
			cfg.MemorystorePassword,
			cfg.MemorystoreHost,
			cfg.MemorystorePort,
			cfg.MemorystoreIndex)
	default:
		return fmt.Sprintf("redis://%s:%d/%d",
			cfg.MemorystoreHost,
			cfg.MemorystorePort,
			cfg.MemorystoreIndex)
	}
}
