package main

import (
	"fmt"
	"log"
	"time"

	"go.mod/cache"
)

func main() {
	fmt.Println("Hello")

	cachestorage := cache.New()
	fmt.Printf("Пустой кэш:%v\n", cachestorage.Items)

	for i := 0; i < 1000; i++ {
		cachestorage.Set(fmt.Sprintf("KEY%d", i), fmt.Sprintf("VALUE%d", i), time.Second*time.Duration(i+1))
	}

	fmt.Printf("Кол-во значений в КЭШ:%v\n", len(cachestorage.Items))

	key, err := cachestorage.Get("KEY0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(key)

	time.Sleep(time.Second * 2)

	for i := 0; i < 1000; i++ {
		key, err = cachestorage.Get(fmt.Sprintf("KEY%d", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Кол-во значений в КЭШ:%v\n", len(cachestorage.Items))

	// fmt.Println(key)

	// fmt.Printf("%v\n", cachestorage)
}
