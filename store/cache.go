package store

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

var cacheStore *cache.Cache

func init() {
	cacheStore = cache.New(time.Hour*24, time.Hour*24)
	fmt.Println("Initializing cacheSTore")
}

// Cache implements the Writer interface
type Cache struct{}

func (c Cache) Write(key string, value interface{}) error {
	return cacheStore.Add(key, value, cache.NoExpiration)
}

func (c Cache) Read(key string) (interface{}, error) {
	value, found := cacheStore.Get(key)
	if !found {
		return "", fmt.Errorf("CacheError: \"%s\" not found in the cache", key)
	}
	cacheStore.Delete(key)
	return value, nil
}
