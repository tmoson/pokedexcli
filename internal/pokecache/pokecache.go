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
	interval time.Duration
	entries  map[string]cacheEntry
	mu       sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.entries[key]
	if exists {
		return entry.val, true
	}
	return []byte{}, false
}

func (c *Cache) reapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.entries {
		if time.Since(entry.createdAt) > c.interval {
			delete(c.entries, key)
		}
	}
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		interval: interval,
		entries:  make(map[string]cacheEntry),
	}
	go func() {
		cleaner := time.NewTicker(interval)
		for range cleaner.C {
			cache.reapLoop()
		}
	}()
	return cache
}
