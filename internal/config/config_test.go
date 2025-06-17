package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func resetEnvVars() {
	os.Setenv("SPACE_APPLICATION_KEY", "application_key")
	os.Setenv("SPACE_DATASTORE_HOST", "localhost")
	os.Setenv("SPACE_DATASTORE_PORT", "5432")
	os.Setenv("SPACE_DATASTORE_NAME_PREFIX", "space")
	os.Setenv("SPACE_DATASTORE_USER", "user")
	os.Setenv("SPACE_DATASTORE_PASSWORD", "password")
	os.Setenv("SPACE_DATASTORE_SSL_MODE", "disable")
	os.Setenv("SPACE_MAIL_FROM", "example@example.com")
	os.Setenv("SPACE_MAILER_ACCESS", "AccessKeyId:SecretAccessKey:Region")
	os.Setenv("SPACE_MEMORY_STORE_HOST", "localhost")
	os.Setenv("SPACE_MEMORY_STORE_PORT", "6379")
	os.Setenv("SPACE_MEMORY_STORE_INDEX", "0")
	os.Setenv("SPACE_MEMORY_STORE_PASSWORD", "")
	os.Setenv("SPACE_SESSION_SECRET", "session_secret")
	os.Setenv("SPACE_SESSION_SECURE", "false")
	os.Setenv("SPACE_STORAGE_SECRET", "storage_secret")
	os.Setenv("SPACE_SENTRY_URL", "")
}

func TestEnvironment(t *testing.T) {
	var cfg *Config
	var err error

	os.Unsetenv("SPACE_ENV")
	resetEnvVars()
	cfg, err = Load()
	assert.Nil(t, err)
	assert.Equal(t, "development", cfg.Environment, "environment is development")
	assert.True(t, cfg.IsEnvironment("development"), "environment is development")
	assert.False(t, cfg.IsEnvironment("production"), "production is not the defined environment")
	assert.False(t, cfg.IsEnvironment("testing"), "testing is not the defined environment")
	os.Setenv("SPACE_ENV", "test")
	resetEnvVars()
	cfg, err = Load()
	assert.Nil(t, err)
	assert.Equal(t, "test", cfg.Environment, "environment is testing")
	assert.True(t, cfg.IsEnvironment("test"), "environment is testing")
	assert.False(t, cfg.IsEnvironment("production"), "production is not the defined environment")
	assert.False(t, cfg.IsEnvironment("development"), "development is not the defined environment")
	os.Setenv("SPACE_ENV", "integration")
	resetEnvVars()
	cfg, err = Load()
	assert.Nil(t, err)
	assert.Equal(t, "integration", cfg.Environment, "environment is testing")
	assert.True(t, cfg.IsEnvironment("integration"), "environment is testing")
	assert.False(t, cfg.IsEnvironment("production"), "production is not the defined environment")
	assert.False(t, cfg.IsEnvironment("development"), "development is not the defined environment")
	os.Setenv("SPACE_ENV", "production")
	resetEnvVars()
	cfg, err = Load()
	assert.Nil(t, err)
	assert.Equal(t, "production", cfg.Environment, "environment is production")
	assert.True(t, cfg.IsEnvironment("production"), "environment is production")
	assert.False(t, cfg.IsEnvironment("testing"), "testing is not the defined environment")
	assert.False(t, cfg.IsEnvironment("development"), "development is not the defined environment")
	os.Setenv("SPACE_ENV", "invalid")
	resetEnvVars()
	cfg, err = Load()
	assert.Nil(t, err)
	assert.Equal(t, "development", cfg.Environment, "development is the default environment for invalid values")
	assert.False(t, cfg.IsEnvironment("production"), "production is not the defined environment")
	assert.False(t, cfg.IsEnvironment("testing"), "testing is not the defined environment")
	assert.True(t, cfg.IsEnvironment("development"), "development is the defined environment")
}
