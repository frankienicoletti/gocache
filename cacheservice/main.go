package cacheservice

import (
	"encoding/json"
	"fmt"
)

// Example struct
type Example struct {
	Name  string
	Count int
}

// GetExampleCachedResults gets caching results and unmarshalls into stype model.Example
func (cacheService *CacheService) GetExampleCachedResults(key string) (cache Example, err error) {
	var data []byte

	// check cache
	if data, err = cacheService.RedisService.GetCache(key); err != nil {
		return cache, fmt.Errorf("Error unable to get key %v from cache: %v", key, err)
	}

	// convert results into model.Example
	if err = json.Unmarshal(data, &cache); err != nil {
		return cache, fmt.Errorf("Error unable to unmarshal cache data: %v", err)
	}

	return cache, nil
}

// CacheResults caches all results
func (cacheService *CacheService) CacheResults(key string, results interface{}) (err error) {
	if err = cacheService.RedisService.CacheResults(key, results, 3600); err != nil {
		return fmt.Errorf("Error unable to cache results: %v", err)
	}

	return nil
}

// ClearCache clears the entire cache
func (cacheService *CacheService) ClearCache() (err error) {
	if err = cacheService.RedisService.ClearAll(); err != nil {
		return fmt.Errorf("Error unable to clear cache: %v", err)
	}

	return nil
}
