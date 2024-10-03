package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	mu       *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cache:    map[string]cacheEntry{},
		mu:       &sync.Mutex{},
		interval: interval,
	}
	go cache.reapLoop()
	return cache
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cache[key] = entry
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.cache[key]
	if ok {
		return val.val, true
	}
	return nil, false
}

func (c Cache) reapLoop() {
	for {
		c.mu.Lock()
		for key, entry := range c.cache {

			if time.Since(entry.createdAt) > c.interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
		time.Sleep(c.interval)
	}
}
