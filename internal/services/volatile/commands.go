package volatile

import (
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
