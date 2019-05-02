package feature

import (
    "github.com/garyburd/redigo/redis"
    "github.com/earaujoassis/space/memstore"
)

// Active is used to check if a feature-gate `name` is currently active (through Redis keys)
func IsActive(name string) bool {
    memstore.Start()
    defer memstore.Close()
    if featureExists, _ := redis.Bool(memstore.Do("HEXISTS", "feature.gates", name)); !featureExists {
        return false
    }
    return true
}

func Enable(name string) {
    memstore.Start()
    defer memstore.Close()
    memstore.Do("HSET", "feature.gates", name, 1)
}

func Disable(name string) {
    memstore.Start()
    defer memstore.Close()
    memstore.Do("HDEL", "feature.gates", name)
}
