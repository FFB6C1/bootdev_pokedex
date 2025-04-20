package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cache: map[string]cacheEntry{},
	}
	go cache.reapLoop(interval)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cache[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if item, ok := c.cache[key]; ok {
		return item.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)

	for {
		_ = <-ticker.C
		c.mu.Lock()
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > duration {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}
