package memory

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type Commands struct {
	conn redis.Conn
}

func NewCommand(conn redis.Conn) *Commands {
	return &Commands{
		conn: conn,
	}
}

func (c *Commands) AddToSortedSetAtKey(key string, score, member interface{}) {
	c.conn.Do("ZADD", key, score, member)
}

func (c *Commands) RemoveFromSortedSetAtKey(key string, member interface{}) {
	c.conn.Do("ZREM", key, member)
}

func (c *Commands) CheckFieldExistence(key, field string) bool {
	keyExists, _ := Bool(c.conn.Do("HEXISTS", key, field))
	return keyExists
}

func (c *Commands) SetFieldAtKey(key, field string, value interface{}) {
	c.conn.Do("HSET", key, field, value)
}

func (c *Commands) SetKeyNXWithExpiration(key string, value interface{}, ttl time.Duration) bool {
	ttlSeconds := int64(ttl.Seconds())
	_, err := c.conn.Do("SET", key, value, "NX", "EX", ttlSeconds)
	return err == nil
}

func (c *Commands) SetKeyWithExpiration(key string, value interface{}, ttl time.Duration) bool {
	ttlSeconds := int64(ttl.Seconds())
	_, err := c.conn.Do("SET", key, value, "EX", ttlSeconds)
	return err == nil
}

func (c *Commands) GetKey(key string) Value {
	result, err := c.conn.Do("GET", key)
	return Value{Result: result, Error: err}
}

func (c *Commands) IncrementFieldAtKeyBy(key, field string, value interface{}) {
	c.conn.Do("HINCRBY", key, field, value)
}

func (c *Commands) GetFieldAtKey(key, field string) Value {
	result, err := c.conn.Do("HGET", key, field)
	return Value{Result: result, Error: err}
}

func (c *Commands) DeleteFieldAtKey(key, field string) {
	c.conn.Do("HDEL", key, field)
}
