package config

import (
    "os"
    "strings"
)

var environment string

func init() {
    environment = strings.ToLower(os.Getenv("ENV"))
    if environment == "" {
        environment = "development"
    }
}

// Environment returns the current environment for the application;
//      it could be: development, testing, production
func Environment() string {
    return environment
}

// IsEnvironment checks if the current environment for the application
//      is the same as defined in `env`
func IsEnvironment(env string) bool {
    return strings.ToLower(env) == Environment()
}

// GetConfig gets a `key` from the environment variables
func GetConfig(key string) string {
    return os.Getenv(key)
}
