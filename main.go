package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type cachestorage struct {
	items map[string]item
}

func newCache() *cachestorage {
	cache := &cachestorage{
		items: make(map[string]item),
	}
	return cache
}

func (c *cachestorage) set(key string, val interface{}, ttl time.Duration) {
	v := newItem(val, ttl)
	c.items[key] = v
}

func (c cachestorage) get(key string) (interface{}, error) {
	c.searchingExpiredValue(key)
	val, exists := c.items[key]
	if !exists {
		return nil, errors.New("key is not exist")
	} else {
		return val.value, nil
	}
}

func (c *cachestorage) delete(key string) error {
	delete(c.items, key)
	return nil
}

func (c *cachestorage) searchingExpiredValue(key string) error {
	t0 := time.Now().Unix()
	for key, item := range c.items {
		if t0 > item.ttl {
			c.delete(key)
		}
	}
	return nil
}

type item struct {
	value interface{}
	ttl   int64
}

func newItem(value interface{}, ttl time.Duration) item {
	item := item{
		value: value,
		ttl:   time.Now().Add(ttl).Unix(),
	}
	return item
}

func main() {
	fmt.Println("Hello")

	cachestorage1 := newCache()
	fmt.Printf("Пустой кэш:%v\n", cachestorage1.items)

	cachestorage1.set("KEY1", "Enythings", time.Second*2)
	cachestorage1.set("KEY2", "TODAY", time.Second*7)
	cachestorage1.set("KEY3", "YESTODAY", time.Second*10)

	fmt.Printf("Кэш значений:%v\n", cachestorage1.items)

	key, err := cachestorage1.get("KEY1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(key)

	time.Sleep(time.Second * 8)

	key, err = cachestorage1.get("KEY2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(key)

	fmt.Printf("%v\n", cachestorage1)
}
