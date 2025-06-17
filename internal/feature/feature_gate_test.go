package feature

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/gateways/redis"
)

func getFeatureGate() *FeatureGate {
	cfg := &config.Config{
		MemorystoreHost:     "localhost",
		MemorystoreIndex:    0,
		MemorystorePassword: "",
		MemorystorePort:     6379,
	}
	ms, _ := redis.NewMemoryService(cfg)
	fg := NewFeatureGate(ms)
	return fg
}

func TestIsActive(t *testing.T) {
	fg := getFeatureGate()
	assert.False(t, fg.IsActive("no-feature"), "shouldn't have no-feature active")
}

func TestEnable(t *testing.T) {
	fg := getFeatureGate()
	assert.False(t, fg.IsActive("not-enabled"), "shouldn't have not-enabled active")
	fg.Enable("not-enabled")
	assert.True(t, fg.IsActive("not-enabled"), "should have no-feature active")
	fg.Disable("not-enabled")
}

func TestDisable(t *testing.T) {
	fg := getFeatureGate()
	assert.False(t, fg.IsActive("to-disable"), "shouldn't have to-disable active")
	fg.Enable("to-disable")
	assert.True(t, fg.IsActive("to-disable"), "should have to-disable active")
	fg.Disable("to-disable")
	assert.False(t, fg.IsActive("to-disable"), "shouldn't have to-disable active")
}
