package cache

import (
	"time"
)

type Cache struct {
	Items map[string]value
}

type value struct {
	item interface{}
	ttl  time.Duration
}

// Make new Cache
func New() *Cache {
	return &Cache{
		Items: make(map[string]value),
	}
}

// Add cache-value to Cache
func (c *Cache) Set(key string, val value, ttl time.Duration) {
	c.Items[key] = val
}

// Get cache-value from Cache
func (c Cache) Get(k string) interface{} {
	value, exists := c.Items[k]
	if exists {
		return value
	} else {
		return nil
	}

}

// Delete cache-value from Cache
func (c *Cache) Delete(k string) error {
	delete(c.Items, k)
	return nil
}
