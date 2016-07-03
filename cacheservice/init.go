package cacheservice

import RedisService "github.com/katejefferson/gocache/redisservice"

//CacheService used to interact with our Cache
type CacheService struct {
	RedisService RedisService.RedisService
}

//Signature used for mocking service
type Signature interface {
	GetExampleCachedResults(key string) (Example, error)
	CacheResults(key string, results interface{}) error
	ClearCache() error
}

//New Constructor for CacheService
func New(connectionString string) CacheService {
	redisService := RedisService.New(connectionString)
	return CacheService{RedisService: redisService}
}
