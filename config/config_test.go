package config

import (
    "testing"
    "os"

    "github.com/stretchr/testify/assert"
)

func TestEnvironment(t *testing.T) {
    externalEnvironment := os.Getenv("ENV")
    if externalEnvironment == "" {
        assert.Equal(t, "", GetConfig("ENV"), "empty externalEnvironment")
        assert.Equal(t, "development", Environment(), "default environment is development")
        assert.True(t, IsEnvironment("development"), "default environment is development")
        assert.False(t, IsEnvironment("production"), "production is not the default environment")
        assert.False(t, IsEnvironment("testing"), "testing is not the default environment")
    } else {
        assert.Equal(t, externalEnvironment, Environment(), "should equal externalEnvironment")
        assert.True(t, IsEnvironment(externalEnvironment), "should return true for the externalEnvironment")
    }
}
