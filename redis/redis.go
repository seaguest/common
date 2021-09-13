package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func GetRedisPool(address, password string, maxConnection int) *redis.Pool {
	pool := &redis.Pool{
		MaxIdle:     maxConnection,
		MaxActive:   maxConnection,
		Wait:        false,
		IdleTimeout: 240 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			return dial("tcp", address, password)
		},
	}
	return pool
}

func dial(network, address, password string) (redis.Conn, error) {
	c, err := redis.Dial(network, address)
	if err != nil {
		return nil, err
	}
	if password != "" {
		if _, err := c.Do("AUTH", password); err != nil {
			c.Close()
			return nil, err
		}
	}
	return c, err
}

func Test(address, password string) error {
	c, err := redis.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer c.Close()
	if password != "" {
		if _, err := c.Do("AUTH", password); err != nil {
			c.Close()
			return err
		}
	}
	_, err = c.Do("PING")
	return err
}

func Incr(key string, ttl int, pool *redis.Pool) error {
	conn := pool.Get()
	defer conn.Close()

	if err := conn.Err(); err != nil {
		return err
	}

	_, err := conn.Do("INCR", key)
	if err != nil {
		return err
	}
	// ttl>0时认为需要设置过期时间
	if ttl > 0 {
		_, err = conn.Do("EXPIRE", key, ttl)
		return err
	}
	return nil
}

func IncrBy(key string, incrAmount uint64, ttl int, pool *redis.Pool) error {
	conn := pool.Get()
	defer conn.Close()

	if err := conn.Err(); err != nil {
		return err
	}

	_, err := conn.Do("INCRBY", key, incrAmount)
	if err != nil {
		return err
	}
	// ttl>0时认为需要设置过期时间
	if ttl > 0 {
		_, err = conn.Do("EXPIRE", key, ttl)
		return err
	}
	return nil
}

func GetKeysByPattern(pattern string, pool *redis.Pool) ([]string, error) {
	conn := pool.Get()
	defer conn.Close()

	if err := conn.Err(); err != nil {
		return []string{}, err
	}

	data, err := redis.Strings(conn.Do("keys", pattern))
	return data, err
}

func SetUint64(key string, value uint64, ttl int, pool *redis.Pool) error {
	conn := pool.Get()
	defer conn.Close()

	if err := conn.Err(); err != nil {
		return err
	}
	_, err := conn.Do("SETEX", key, ttl, value)
	return err

}

func SetKeyExpire(key string, ttl int, pool *redis.Pool) error {
	conn := pool.Get()
	defer conn.Close()

	if err := conn.Err(); err != nil {
		return err
	}
	_, err := conn.Do("EXPIRE", key, ttl)
	if err != nil {
		return err
	}
	return nil
}

func GetUint64(key string, pool *redis.Pool) (uint64, error) {
	conn := pool.Get()
	defer conn.Close()

	if err := conn.Err(); err != nil {
		return 0, err
	}

	n, err := redis.Uint64(conn.Do("GET", key))
	if err != nil {
		return 0, err
	}
	return n, err
}

func SetObject(key string, obj interface{}, ttl int, pool *redis.Pool) error {
	conn := pool.Get()
	defer conn.Close()

	bs, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	if err := conn.Err(); err != nil {
		return err
	}
	_, err = conn.Do("SETEX", key, ttl, string(bs))
	return err
}

func GetObject(key string, obj interface{}, pool *redis.Pool) (bool, error) {
	conn := pool.Get()
	defer conn.Close()

	if err := conn.Err(); err != nil {
		return false, err
	}

	s, err := redis.String(conn.Do("GET", key))
	if err == redis.ErrNil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	err = json.Unmarshal([]byte(s), obj)
	if err != nil {
		return false, err
	}
	return true, nil
}

func SetString(key, value string, ttl int, pool *redis.Pool) error {
	conn := pool.Get()
	defer conn.Close()

	if err := conn.Err(); err != nil {
		return err
	}
	_, err := conn.Do("SETEX", key, ttl, value)
	return err
}

func GetString(key string, pool *redis.Pool) (string, error) {
	conn := pool.Get()
	defer conn.Close()

	if err := conn.Err(); err != nil {
		return "", err
	}

	s, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}
	return s, nil
}

func SetInt(key string, value, ttl int, pool *redis.Pool) error {
	conn := pool.Get()
	defer conn.Close()

	if err := conn.Err(); err != nil {
		return err
	}
	_, err := conn.Do("SETEX", key, ttl, value)
	return err
}

func GetInt(key string, pool *redis.Pool) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	if err := conn.Err(); err != nil {
		return 0, err
	}

	n, err := redis.Int(conn.Do("GET", key))
	if err != nil {
		return 0, err
	}
	return n, err
}

func DelKey(key string, pool *redis.Pool) error {
	conn := pool.Get()
	defer conn.Close()

	var err error
	if err = conn.Err(); err != nil {
		return err
	}

	_, err = conn.Do("DEL", key)
	return err
}

func Publish(channel, message string, pool *redis.Pool) error {
	conn := pool.Get()
	defer conn.Close()

	if err := conn.Err(); err != nil {
		return err
	}

	_, err := conn.Do("PUBLISH", channel, message)
	if err != nil {
		return err
	}
	return nil
}

func GetSetUint64(key string, value uint64, pool *redis.Pool) (uint64, error) {
	conn := pool.Get()
	defer conn.Close()

	if err := conn.Err(); err != nil {
		return 0, err
	}

	n, err := redis.Uint64(conn.Do("GETSET", key, value))
	if err != nil {
		return 0, err
	}
	return n, err
}
