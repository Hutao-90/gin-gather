package cachepackage

import (
	"github.com/patrickmn/go-cache"
)

var (
	cacheInstance *cache.Cache
)

func init() {
	// Initialize the cache instance
	cacheInstance = cache.New(cache.NoExpiration, cache.DefaultExpiration)
}

func GetCacheInstance() *cache.Cache {
	// Return the cache instance
	return cacheInstance
}
