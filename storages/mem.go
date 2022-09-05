package storages

import (
	"gitlab.com/danmory/web-hashing-server/tools"
)

type memoryStorage struct {
	data map[string]string
}

func createMemStorage() *memoryStorage {
	memStorage := new(memoryStorage)
	memStorage.data = make(map[string]string)
	return memStorage
}

func (mstor *memoryStorage) Store(value string) (string, error) {
	if !tools.IsURL(value) {
		return "", &storageError{reason: "The value " + value + " is not URL"}
	}
	key := tools.StringConverter.Do(value)
	if _, exists := mstor.data[key]; exists {
		return "", &storageError{reason: "The value " + value + " is already added"}
	}
	mstor.data[key] = value
	return key, nil
}

func (mstor *memoryStorage) Find(key string) (string, error) {
	value, exists := mstor.data[key]
	if !exists {
		return "", &storageError{reason: "The key " + key + " does not exist"}
	}
	return value, nil
}
