package redis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/logs"
)

type MemoryService struct {
	pool   *redis.Pool
	config *config.Config
}

func NewMemoryService(cfg *config.Config) (*MemoryService, error) {
	pool := &redis.Pool{
		MaxIdle:     10,
		MaxActive:   100,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(getRedisURL(cfg))
			if err != nil {
				logs.Propagate(logs.Error, err.Error())
				return nil, err
			}

			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				logs.Propagate(logs.Error, err.Error())
			}
			return err
		},
	}

	conn := pool.Get()
	defer conn.Close()
	if _, err := conn.Do("PING"); err != nil {
		logs.Propagate(logs.Error, err.Error())
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return &MemoryService{
		pool:   pool,
		config: cfg,
	}, nil
}

func (ms *MemoryService) GetPool() *redis.Pool {
	return ms.pool
}

func (ms *MemoryService) Transaction(transaction func(*Commands)) {
	conn := ms.pool.Get()
	defer conn.Close()

	commands := NewCommand(conn)
	transaction(commands)
}

func (ms *MemoryService) Do(commandName string, args ...interface{}) (interface{}, error) {
	conn := ms.pool.Get()
	defer conn.Close()
	return conn.Do(commandName, args...)
}

func (ms *MemoryService) Close() error {
	return ms.pool.Close()
}
