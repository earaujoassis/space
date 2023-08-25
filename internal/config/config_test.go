package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironment(t *testing.T) {
	externalEnvironment := os.Getenv("SPACE_ENV")
	assert.Equal(t, "", GetEnvVar("UNSET_VAR"), "empty UNSET_VAR")
	if externalEnvironment == "" {
		assert.Equal(t, "development", Environment(), "default environment is development")
		assert.True(t, IsEnvironment("development"), "default environment is development")
		assert.False(t, IsEnvironment("production"), "production is not the default environment")
		assert.False(t, IsEnvironment("testing"), "testing is not the default environment")
	} else if externalEnvironment == "testing" {
		assert.Equal(t, "testing", Environment(), "set environment is testing")
		assert.True(t, IsEnvironment("testing"), "set environment is testing")
		assert.False(t, IsEnvironment("production"), "production is not the set environment")
		assert.False(t, IsEnvironment("development"), "development is not the set environment")
	} else {
		assert.Equal(t, externalEnvironment, Environment(), "should equal externalEnvironment")
		assert.True(t, IsEnvironment(externalEnvironment), "should return true for the externalEnvironment")
	}
}
