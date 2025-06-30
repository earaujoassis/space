package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/logs"
)

type MemoryService struct {
	pool   *redis.Pool
	config *config.Config
}

func NewMemoryService(cfg *config.Config) (*MemoryService, error) {
	pool := NewRedisProviderPool(cfg)
	if pool == nil {
		return nil, fmt.Errorf("failed to create pool from provider")
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
