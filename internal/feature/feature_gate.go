package feature

import (
	"github.com/earaujoassis/space/internal/gateways/redis"
)

type FeatureGate struct {
	ms *redis.MemoryService
}

func NewFeatureGate(ms *redis.MemoryService) *FeatureGate {
	return &FeatureGate{
		ms: ms,
	}
}

// IsActive is used to check if a feature-gate `name` is currently active (through Redis keys)
func (fg *FeatureGate) IsActive(name string) bool {
	var result bool

	fg.ms.Transaction(func(c *redis.Commands) {
		result = c.CheckFieldExistence("feature.gates", name)
	})

	return result
}

// Enable makes a feature-gate `name` currently active (through Redis keys)
func (fg *FeatureGate) Enable(name string) {
	fg.ms.Transaction(func(c *redis.Commands) {
		c.SetFieldAtKey("feature.gates", name, 1)
	})
}

// Disable makes a feature-gate `name` currently inactive (through Redis keys)
func (fg *FeatureGate) Disable(name string) {
	fg.ms.Transaction(func(c *redis.Commands) {
		c.DeleteFieldAtKey("feature.gates", name)
	})
}
