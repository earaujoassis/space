package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/logs"
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
		logs.Propagate(logs.Panic, err.Error())
	}
}

// Close is used to end a connection to a Redis datastore; given an active connection
func Close() {
	if memoryStore != nil {
		memoryStore.Close()
		memoryStore = nil
	}
}
