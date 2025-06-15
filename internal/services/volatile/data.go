package volatile

import (
	"github.com/gomodule/redigo/redis"
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

func Bool(reply interface{}, err error) (bool, error) {
	return redis.Bool(reply, err)
}
