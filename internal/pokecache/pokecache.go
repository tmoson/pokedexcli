package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	Interval time.Duration
	Entries  map[string]cacheEntry
	mu       sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Entries[key] = cacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.Entries[key]
	if exists {
		return entry.Val, true
	}
	return []byte{}, false
}

func (c *Cache) reapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.Entries {
		if time.Since(entry.CreatedAt) > c.Interval {
			delete(c.Entries, key)
		}
	}
}

func NewCache(Interval time.Duration) Cache {
	cache := Cache{
		Interval: Interval,
		Entries:  make(map[string]cacheEntry),
	}
	go func() {
		cleaner := time.NewTicker(Interval)
		for range cleaner.C {
			cache.reapLoop()
		}
	}()
	return cache
}
