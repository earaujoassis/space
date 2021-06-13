package memstore

import (
    "fmt"

    "github.com/garyburd/redigo/redis"

    "github.com/earaujoassis/space/internal/config"
)

var memoryStore redis.Conn

// Start is used to setup a connection with a Redis datastore
func Start() {
    var err error
    var storeURI string
    var cfg config.Config = config.GetGlobalConfig()
    if config.IsEnvironment("production") {
        storeURI = fmt.Sprintf("redis://:%v@%v:%v/%v",
            cfg.MemorystorePassword,
            cfg.MemorystoreHost,
            cfg.MemorystorePort,
            cfg.MemorystoreIndex)
    } else {
        storeURI = fmt.Sprintf("redis://%v:%v/%v",
            cfg.MemorystoreHost,
            cfg.MemorystorePort,
            cfg.MemorystoreIndex)
    }
    memoryStore, err = redis.DialURL(storeURI)
    if err != nil {
        panic(err)
    }
}

// Do is used to send commands to a Redis datastore, throguh an active connection
func Do(commandName string, args ...interface{}) (reply interface{}, err error) {
    return memoryStore.Do(commandName, args...)
}

// Close is used to end a connection to a Redis datastore; given an active connection
func Close() {
    if memoryStore != nil {
        memoryStore.Close()
        memoryStore = nil
    }
}
