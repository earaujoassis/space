package repository

import (
	"github.com/earaujoassis/space/internal/gateways/redis"
)

type BaseMemoryRepository[T any] struct {
	ms *redis.MemoryService
}

func NewBaseMemoryRepository[T any](ms *redis.MemoryService) *BaseMemoryRepository[T] {
	return &BaseMemoryRepository[T]{ms: ms}
}
