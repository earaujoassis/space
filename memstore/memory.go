package memstore

import (
    "fmt"

    "github.com/garyburd/redigo/redis"

    "github.com/earaujoassis/space/config"
)

var memoryStore redis.Conn

// Start is used to setup a connection with a Redis datastore
func Start() {
    var err error
    var storeURI string
    if config.IsEnvironment("production") {
        storeURI = fmt.Sprintf("redis://:%v@%v:%v/%v",
            config.GetConfig("SPACE_MEMORYSTORE_PASSWORD"),
            config.GetConfig("SPACE_MEMORYSTORE_HOST"),
            config.GetConfig("SPACE_MEMORYSTORE_PORT"),
            config.GetConfig("SPACE_MEMORYSTORE_INDEX"))
    } else {
        storeURI = fmt.Sprintf("redis://%v:%v/%v",
            config.GetConfig("SPACE_MEMORYSTORE_HOST"),
            config.GetConfig("SPACE_MEMORYSTORE_PORT"),
            config.GetConfig("SPACE_MEMORYSTORE_INDEX"))
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
