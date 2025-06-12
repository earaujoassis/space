package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironment(t *testing.T) {
	os.Unsetenv("SPACE_ENV")
	assert.Equal(t, "", Environment(), "environment is empty")
	os.Setenv("SPACE_ENV", "development")
	LoadEnvironment()
	assert.Equal(t, "development", Environment(), "environment is development")
	assert.True(t, IsEnvironment("development"), "environment is development")
	assert.False(t, IsEnvironment("production"), "production is not the defined environment")
	assert.False(t, IsEnvironment("testing"), "testing is not the defined environment")
	os.Setenv("SPACE_ENV", "testing")
	LoadEnvironment()
	assert.Equal(t, "testing", Environment(), "environment is testing")
	assert.True(t, IsEnvironment("testing"), "environment is testing")
	assert.False(t, IsEnvironment("production"), "production is not the defined environment")
	assert.False(t, IsEnvironment("development"), "development is not the defined environment")
	os.Setenv("SPACE_ENV", "production")
	LoadEnvironment()
	assert.Equal(t, "production", Environment(), "environment is production")
	assert.True(t, IsEnvironment("production"), "environment is production")
	assert.False(t, IsEnvironment("testing"), "testing is not the defined environment")
	assert.False(t, IsEnvironment("development"), "development is not the defined environment")
	os.Setenv("SPACE_ENV", "invalid")
	LoadEnvironment()
	assert.Equal(t, "", Environment(), "environment is not defined")
	assert.False(t, IsEnvironment("production"), "production is not the defined environment")
	assert.False(t, IsEnvironment("testing"), "testing is not the defined environment")
	assert.False(t, IsEnvironment("development"), "development is not the defined environment")
}
