package redisservice

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// GetCache gets previously cached results, returns nil if not found
func (redisService *RedisService) GetCache(key string) (n []byte, err error) {
	// connect to Redis container
	c, err := redisService.getRedisConnection()
	if err != nil {
		return nil, fmt.Errorf("Error unable to get redis connection: %v", err)
	}
	defer c.Close()

	// get cached results from Redis and convert to Byte array
	n, err = redis.Bytes(c.Do("GET", key))
	if err != nil {
		return nil, fmt.Errorf("Error unable to convert redis data to bytes: %v", err)
	}

	return n, nil
}

// CacheResults caches results
func (redisService *RedisService) CacheResults(key string, data interface{}, expiry int) (err error) {
	var jsonBytes []byte

	// convert results to string
	if jsonBytes, err = json.Marshal(data); err != nil {
		return fmt.Errorf("Error unable to marshal cache data: %v", err)
	}
	stringItems := string(jsonBytes)

	// connect to Redis container
	c, err := redisService.getRedisConnection()
	if err != nil {
		return fmt.Errorf("Error unable to get redis connection: %v", err)
	}
	defer c.Close()

	// cache string and expire it
	c.Do("MULTI")
	c.Do("SET", key, stringItems)
	c.Do("EXPIRE", key, expiry)
	_, err = c.Do("EXEC")

	return err
}

// ClearAll clears the entire cache
func (redisService *RedisService) ClearAll() (err error) {

	c, err := redisService.getRedisConnection()
	if err != nil {
		return fmt.Errorf("Error unable to get redis connection: %v", err)
	}
	defer c.Close()

	c.Do("FLUSHALL")

	return nil
}

// getRedisConnection connects to the Redis
func (redisService *RedisService) getRedisConnection() (c redis.Conn, err error) {
	if c, err = redis.Dial("tcp", redisService.ConnectionString); err != nil {
		return nil, fmt.Errorf("Error redis dial failure: %v", err)
	}
	return c, nil
}
