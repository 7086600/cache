package cache

type Cache struct {
	items map[string]interface{}
}

// Make new Cache
func New() *Cache {
	return &Cache{
		items: make(map[string]interface{}),
	}
}

// Add cache-value to Cache
func (c *Cache) Set(key string, value interface{}) {
	c.items[key] = value
}

// Get cache-value from Cache
func (c Cache) Get(k string) interface{} {
	value, exists := c.items[k]
	if exists {
		return value
	} else {
		return nil
	}

}

// Delete cache-value from Cache
func (c *Cache) Delete(k string) {
	delete(c.items, k)
}
