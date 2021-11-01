# cache
build in-memory cache in Go

# Install

go get -u github.com/7086600/cache

# Example

package main

import (
	"fmt"

	"github.com/7086600/cache"
)

func main() {
	// Make new cache
	cache := cache.New()

	// Set value to cache
	cache.Set("Key", []string{"Veni", "Vidi", "Vici"})
	// Set value to cache with for:
	for i := 0; i < 10; i++ {
		cache.Set((fmt.Sprintf("Key%d", i)), (fmt.Sprintf("Value%d", i)))
	}

	// Get info about value of cache
	key1 := cache.Get("Key1")
	fmt.Println(key1)

	// Delete value of cachу
	cache.Delete("Key8")
	key8 := cache.Get("Key8")

	fmt.Println(key8)

}

