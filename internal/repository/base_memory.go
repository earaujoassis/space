package repository

import (
	"github.com/earaujoassis/space/internal/gateways/memory"
)

type BaseMemoryRepository[T any] struct {
	ms *memory.MemoryService
}

func NewBaseMemoryRepository[T any](ms *memory.MemoryService) *BaseMemoryRepository[T] {
	return &BaseMemoryRepository[T]{ms: ms}
}
