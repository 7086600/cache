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
	fmt.Printf("Кол-во значений в КЭШ:%v\n", len(cachestorage.Items))

	cachestorage.Set("userID", 42, time.Second*1)
	userId, err := cachestorage.Get("userID")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userId)

	time.Sleep(time.Second * 3)

	userId, err = cachestorage.Get("userId")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	for i := 0; i < 7000; i++ {
		cachestorage.Set(fmt.Sprintf("KEY%d", i), fmt.Sprintf("VALUE%d", i), time.Second*time.Duration(i+1))
	}
	fmt.Printf("Кол-во значений в КЭШ:%v\n", len(cachestorage.Items))

}
