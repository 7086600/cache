package main

import (
	"cachestorage/cache"
	"fmt"
)

func main() {
	// Создаём объект - хранилище для кэша. Используем функцию-конструктор из пакета cache
	cacheStorage1 := cache.NewCacheStorage(map[string]string{})

	// Записываем в поле CacheStorageField первое значение
	cacheStorage1.SetCacheInfo("Level101", "Low")

	// Определим цикл, который создаёт и записывает значения в кэш
	for i := 0; i < 10; i++ {
		cacheStorage1.SetCacheInfo(fmt.Sprintf("Cache%d", i), fmt.Sprintf("Level%d", i))
	}

	// Удаляем значение из кэша
	cacheStorage1.DeleteCacheInfo("Cache5")
	cacheStorage1.DeleteCacheInfo("Cache3")

	// Записываем новое значение в кэш по ключу "Cache4"
	cacheStorage1.SetCacheInfo("Cache4", "NewLevel4")

	// Выводим значение кэша по ключу
	fmt.Println(cacheStorage1.GetCacheInfo("Cache4"))

	// Определим цикл для вывода значений из кэша на экран
	fmt.Println("There is a cache storage:")
	for key := range cacheStorage1.CacheStorageField {
		fmt.Println(cacheStorage1.GetCacheInfo(key))
	}
	fmt.Println()
	cacheStorage1.SetCacheInfo("CACHE11", "Level11")
	cacheKey := cacheStorage1.GetCacheInfo("CACHE11")

	fmt.Println(cacheKey)

	cacheStorage1.DeleteCacheInfo("Cache6")
	cacheKey6 := cacheStorage1.GetCacheInfo("Cache6")
	fmt.Println(cacheKey6)
	cacheStorage1.SetCacheInfo("CACHE6", "NewLevel6")
	cacheKey6 = cacheStorage1.GetCacheInfo("CACHE6")
	fmt.Println(cacheKey6)

}
