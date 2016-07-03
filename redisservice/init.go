package redisservice

import "github.com/garyburd/redigo/redis"

//RedisService used to interact with Redis Cache
type RedisService struct {
	ConnectionString string
}

//Signature used for mocking service
type Signature interface {
	GetCache(key string) ([]byte, error)
	CacheResults(key string, data interface{}) error
	ClearAll() error
	getRedisConnection() (redis.Conn, error)
}

//New Constructor for RedisService
func New(connectionString string) RedisService {
	return RedisService{ConnectionString: connectionString}
}
