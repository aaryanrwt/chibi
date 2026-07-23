package k8s

import (
	"fmt"
	"time"

	"github.com/dgraph-io/ristretto"
)

var Cache *ristretto.Cache

// InitCache initializes the Ristretto in-memory cache.
func InitCache() error {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		return fmt.Errorf("failed to initialize cache: %w", err)
	}
	Cache = cache
	return nil
}

// GetCachedResource attempts to fetch from cache, otherwise returns nil
func GetCachedResource(key string) (interface{}, bool) {
	if Cache == nil {
		return nil, false
	}
	return Cache.Get(key)
}

// SetCachedResource sets a value in the cache with a cost and TTL
func SetCachedResource(key string, value interface{}, cost int64, ttl time.Duration) {
	if Cache != nil {
		Cache.SetWithTTL(key, value, cost, ttl)
	}
}
