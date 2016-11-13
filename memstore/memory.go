package memstore

import (
    "fmt"

    "github.com/garyburd/redigo/redis"

    "github.com/earaujoassis/space/config"
)

var memoryStore redis.Conn

func Start() {
    var err error
    var storeURI string
    if config.GetConfig("environment") == "production" {
        storeURI = fmt.Sprintf("redis://:%v@%v:%v/%v",
            config.GetConfig("memorystore.password"),
            config.GetConfig("memorystore.host"),
            config.GetConfig("memorystore.port"),
            config.GetConfig("memorystore.index"))
    } else {
        storeURI = fmt.Sprintf("redis://%v:%v/%v",
            config.GetConfig("memorystore.host"),
            config.GetConfig("memorystore.port"),
            config.GetConfig("memorystore.index"))
    }
    memoryStore, err = redis.DialURL(storeURI)
    if err != nil {
        panic(err)
    }
}

func Do(commandName string, args ...interface{}) (reply interface{}, err error) {
    return memoryStore.Do(commandName, args...)
}

func Close() {
    if memoryStore != nil {
        memoryStore.Close()
        memoryStore = nil
    }
}
