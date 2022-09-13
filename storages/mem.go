package storages

import (
	"sync"

	"github.com/danmory/web-hashing-server/tools"
)

type memoryStorage struct {
	data sync.Map // string -> string
}

func createMemStorage() *memoryStorage {
	memStorage := new(memoryStorage)
	return memStorage
}

func (mstor *memoryStorage) Store(value string) (string, error) {
	if !tools.IsURL(value) {
		return "", &storageError{reason: "The value " + value + " is not URL"}
	}
	key := tools.StringConverter.Do(value)
	if _, exists := mstor.data.Load(key); exists {
		return "", &storageError{reason: "The value " + value + " is already added"}
	}
	mstor.data.Store(key, value)
	return key, nil
}

func (mstor *memoryStorage) Find(key string) (string, error) {
	value, exists := mstor.data.Load(key)
	if !exists {
		return "", &storageError{reason: "The key " + key + " does not exist"}
	}
	return value.(string), nil
}

func (mstor *memoryStorage) Close() error {
	return nil
}
