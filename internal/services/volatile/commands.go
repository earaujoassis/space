package volatile

import (
	"time"

	memstore "github.com/earaujoassis/space/internal/gateways/redis"
)

func AddToSortedSetAtKey(key string, score, member interface{}) {
	memstore.Do("ZADD", key, score, member)
}

func RemoveFromSortedSetAtKey(key string, member interface{}) {
	memstore.Do("ZREM", key, member)
}

func CheckFieldExistence(key, field string) bool {
	keyExists, _ := Bool(memstore.Do("HEXISTS", key, field))
	return keyExists
}

func SetFieldAtKey(key, field string, value interface{}) {
	memstore.Do("HSET", key, field, value)
}

func SetKeyNXWithExpiration(key string, value interface{}, ttl time.Duration) bool {
	_, err := memstore.Do("SET", key, value, "NX", "EX", ttl)
	return err == nil
}

func SetKeyWithExpiration(key string, value interface{}, ttl time.Duration) {
	memstore.Do("SET", key, value, "EX", ttl)
}

func GetKey(key string) Value {
	result, err := memstore.Do("GET", key)
	return Value{Result: result, Error: err}
}

func IncrementFieldAtKeyBy(key, field string, value interface{}) {
	memstore.Do("HINCRBY", key, field, value)
}

func GetFieldAtKey(key, field string) Value {
	result, err := memstore.Do("HGET", key, field)
	return Value{Result: result, Error: err}
}

func DeleteFieldAtKey(key, field string) {
	memstore.Do("HDEL", key, field)
}

func TransactionWrapper(f func()) {
	memstore.Start()
	defer memstore.Close()

	f()
}
