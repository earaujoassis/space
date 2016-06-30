package main

import (
    "fmt"
    "os"
    "github.com/spf13/viper"
)

func GetConfig(key string) interface{} {
    var environment = os.Getenv("ENVIRONMENT")
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
    return viper.Get(key)
}
