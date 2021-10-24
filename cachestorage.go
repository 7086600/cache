package cachestorage

import "fmt"

// Определим структуру для создания хранилища кэша
// Одно поле CacheStorageField типа мапа (CacheStorageField["String"] = "String")
type CacheStorage struct {
	CacheStorageField map[string]string
}

// Определим функцию-конструктор для создания экземпляров хранилищ кэша
// Ф-ия принимает мапу, возвращает объект - с инициализированной мапой в
// поле CacheStorageField
func NewCacheStorage(map[string]string) CacheStorage {
	return CacheStorage{
		CacheStorageField: make(map[string]string),
	}
}

// Определим метод добавления информации в кэш для нашей структуры  CacheStorage
// Ресивер в виде указателя на структуру, чтобы вносить изменения в сам объект, а
// не в его копию.
func (sc *CacheStorage) SetCacheInfo(newKey, newValue string) {
	sc.CacheStorageField[newKey] = newValue
}

// Определим метод для чтения информации из кэша
// Метод принимает в качестве аргумента ключ - k и возвращает строку: ключ | значение
func (gc CacheStorage) GetCacheInfo(k string) string {
	value, exists := gc.CacheStorageField[k]
	if exists {
		return fmt.Sprintf("Key: %s | Value: %s", k, value)
	} else {
		return fmt.Sprintf("Key: %s | Value doesn't exist", k)
	}

}

// Определим метод для удаления информации из кэша по ключу
func (dc *CacheStorage) DeleteCacheInfo(k string) {
	delete(dc.CacheStorageField, k)
}
