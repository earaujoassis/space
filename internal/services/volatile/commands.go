package volatile

import (
	"github.com/gomodule/redigo/redis"

	"github.com/earaujoassis/space/internal/memstore"
)

type Value struct {
	Result interface{}
	Error  error
}

func (value Value) ToInt64() int64 {
	result, _ := redis.Int64(value.Result, value.Error)
	return result
}

func (value Value) ToInt() int {
	result, _ := redis.Int(value.Result, value.Error)
	return result
}

func (value Value) ToString() string {
	result, _ := redis.String(value.Result, value.Error)
	return result
}

func AddToSortedSetAtKey(key string, score, member interface{}) {
	memstore.Do("ZADD", key, score, member)
}

func RemoveFromSortedSetAtKey(key string, member interface{}) {
	memstore.Do("ZREM", key, member)
}

func CheckFieldExistence(key, field string) bool {
	keyExists, _ := redis.Bool(memstore.Do("HEXISTS", key, field))
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

func TransactionsWrapper(f func()) {
	memstore.Start()
	defer memstore.Close()

	f()
}
