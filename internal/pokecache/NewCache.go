package pokecache

import "time"

func NewCache(interval time.Duration) *Cache {
  // init values
	cache := &Cache{}
	cache.entries = map[string]cacheEntry{}

  // start the cache reapLoop
	cache.reapLoop(interval)
	return cache
}
