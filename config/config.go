package config

import (
    "os"
    "strings"
)

var environment string

func init() {
    environment = strings.ToLower(os.Getenv("ENVIRONMENT"))
    if environment == "" {
        environment = "development"
    }
}

func Environment() string {
    return environment
}

func IsEnvironment(env string) bool {
    return strings.ToLower(env) == Environment()
}

func GetConfig(key string) string {
    return os.Getenv(key)
}
