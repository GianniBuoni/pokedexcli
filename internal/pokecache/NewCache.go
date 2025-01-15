package pokecache

import "time"

func NewCahe(interval time.Duration) {
	var cache Cache
	cache.reapLoop(interval)
}
