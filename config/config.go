package config

import (
    "fmt"
    "os"
    "strings"
    "github.com/spf13/viper"
)

func init() {
    var environment = strings.ToLower(os.Getenv("ENVIRONMENT"))
    if environment == "" {
        environment = "development"
    }
    viper.SetDefault("environment", environment)
    viper.SetConfigName(environment)
    viper.AddConfigPath("../config")
    viper.AddConfigPath("./config")
    err := viper.ReadInConfig()
    if err != nil {
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }
}

func Environment() string {
    return viper.Get("environment").(string)
}

func IsEnvironment(env string) bool {
    return strings.ToLower(env) == Environment()
}

func GetConfig(key string) interface{} {
    return viper.Get(key)
}

func SetConfig(key, value string) {
    viper.SetDefault(key, value)
}
