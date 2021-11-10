# cache
build in-memory cache in Go with ttl value

# Install

go get -u github.com/7086600/cache

# Example

```
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/7086600/cache"
)

func main() {
	// Make new cache
	cache := cache.New()

	// Set value to cache
	cache.Set("Key", []string{"Veni", "Vidi", "Vici"}, time.Second*5)

	// Get value from cache and check errors
	key, err := cache.Get("Key")
	fmt.Println(len(cache.Items))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(key)

	// Pause
	time.Sleep(time.Second * 6)

	key, err = cache.Get("Key")

	key, _ = cache.Get("Key")
	fmt.Println(len(cache.Items))
	if err != nil {
		log.Fatal(err)
	}

}
```