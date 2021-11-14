package cache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	mu    *sync.RWMutex
	Items map[string]item
}

type item struct {
	value interface{}
	ttl   int64
}

// Make new Cache
func New() *Cache {
	cache := &Cache{
		mu:    new(sync.RWMutex),
		Items: make(map[string]item),
	}
	go cache.scanCache()
	return cache
}

// new item of cache
func newItem(value interface{}, ttl time.Duration) item {
	return item{
		value: value,
		ttl:   time.Now().Add(ttl).Unix(),
	}
}

// Add cache-value to Cache
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	c.Items[key] = newItem(value, ttl)
	c.mu.Unlock()
}

// Get cache-value from Cache
func (c Cache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, exists := c.Items[key]
	if !exists {
		return nil, errors.New("")
	}
	return item.value, nil
}

// Delete cache-value from Cache
func (c *Cache) Delete(key string) error {
	delete(c.Items, key)
	return nil
}

// Seacrhing expired value method
func (c *Cache) searchingExpiredValue() error {
	c.mu.Lock()
	for key, item := range c.Items {
		if time.Now().Unix() > item.ttl {
			c.Delete(key)
		}
	}
	c.mu.Unlock()
	return nil
}

// Scanning cache for expired value
func (c *Cache) scanCache() {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		c.searchingExpiredValue()
	}

}
