package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu      sync.Mutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()

	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.entries[key] = newEntry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (data []byte, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			for k, v := range c.entries {
				if time.Since(v.createdAt) > interval {
					c.mu.Lock()
					delete(c.entries, k)
					c.mu.Unlock()
				}
			}
		}
	}()

}
