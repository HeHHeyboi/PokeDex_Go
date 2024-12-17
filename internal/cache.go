package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mu    *sync.Mutex
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		make(map[string]cacheEntry),
		&sync.Mutex{},
	}
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			newCache.reapLoop(interval)

		}
	}()
	return &newCache
}
func (c *Cache) Add(key string, val []byte) {
	cache := cacheEntry{time.Now(), val}
	fmt.Println(cache)
	c.entry[key] = cache
}
func (c *Cache) Get(key string) ([]byte, bool) {
	result, ok := c.entry[key]
	if !ok {
		return nil, false
	}
	return result.val, true
}
func (c *Cache) reapLoop(interval time.Duration) {
	now := time.Now()
	for k, v := range c.entry {
		if now.Sub(v.createdAt) > interval {
			delete(c.entry, k)
		}
	}

}
