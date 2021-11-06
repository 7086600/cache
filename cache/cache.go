package cache

import (
	"errors"
	"time"
)

type Cache struct {
	Items map[string]item
}

type item struct {
	value interface{}
	ttl   int64
}

// Make new Cache
func New() *Cache {
	return &Cache{
		Items: make(map[string]item),
	}
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
	c.Items[key] = newItem(value, ttl)
}

// Get cache-value from Cache
func (c Cache) Get(key string) (interface{}, error) {
	c.searchingExpiredValue(key)
	item, exists := c.Items[key]
	if !exists {
		return nil, errors.New("key is not exist")
	}
	return item.value, nil
}

// Delete cache-value from Cache
func (c *Cache) Delete(key string) error {
	delete(c.Items, key)
	return nil
}

func (c *Cache) searchingExpiredValue(key string) error {
	t0 := time.Now().Unix()
	for key, item := range c.Items {
		if t0 > item.ttl {
			c.Delete(key)
		}
	}
	return nil
}
