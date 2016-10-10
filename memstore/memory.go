package memstore

import (
    "fmt"

    "github.com/garyburd/redigo/redis"

    "github.com/earaujoassis/space/config"
)

var memoryStore redis.Conn

func Start() {
    var err error
    var storeName string = fmt.Sprintf("%v%v",
        config.GetConfig("memorystore.url"),
        config.GetConfig("memorystore.index"))
    memoryStore, err = redis.DialURL(storeName)
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
