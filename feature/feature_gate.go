package feature

import (
    "github.com/garyburd/redigo/redis"
    "github.com/earaujoassis/space/memstore"
)

// Active is used to check if a feature-gate `name` is currently active (through Redis keys)
func Active(name string) bool {
    memstore.Start()
    defer memstore.Close()
    if featureExists, _ := redis.Bool(memstore.Do("HEXISTS", "feature.gates", name)); !featureExists {
        return false
    }
    return true
}
