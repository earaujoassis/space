package config

import (
    "testing"
    "os"
    "strings"

    "github.com/stretchr/testify/assert"
)

func TestEnvironment(t *testing.T) {
    externalEnvironment := strings.ToLower(os.Getenv("ENVIRONMENT"))
    if externalEnvironment == "" {
        assert.Equal(t, "development", Environment(), "default environment is development")
        assert.True(t, IsEnvironment("development"), "default environment is development")
        assert.False(t, IsEnvironment("production"), "production is not the default environment")
        assert.False(t, IsEnvironment("testing"), "testing is not the default environment")
        assert.Equal(t, "development", GetConfig("environment").(string), "default environment is development")
    } else {
        assert.Equal(t, externalEnvironment, Environment(), "should equal externalEnvironment")
        assert.True(t, IsEnvironment(externalEnvironment), "should return true for the externalEnvironment")
        assert.Equal(t, externalEnvironment, GetConfig("environment").(string), "should equal externalEnvironment")
    }
}
