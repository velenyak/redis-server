package storage

import "sync"

var stringStorage = make(map[string]string)
var stringStorageMUs = sync.RWMutex{}

func Set(key string, value string) error {
	stringStorageMUs.Lock()
	stringStorage[key] = value
	stringStorageMUs.Unlock()
	return nil
}

func Get(key string) (string, error) {
	stringStorageMUs.RLock()
	value := stringStorage[key]
	stringStorageMUs.RUnlock()
	return value, nil
}

var hashStorage = make(map[string]map[string]string)
var hashStorageMUs = sync.RWMutex{}

func HSet(key string, field string, value string) error {
	hashStorageMUs.Lock()
	if hashStorage[key] == nil {
		hashStorage[key] = make(map[string]string)
	}
	hashStorage[key][field] = value
	hashStorageMUs.Unlock()
	return nil
}

func HGet(key string, field string) (string, error) {
	hashStorageMUs.RLock()
	value := hashStorage[key][field]
	hashStorageMUs.RUnlock()
	return value, nil
}

func HGetAll(key string) (map[string]string, error) {
	hashStorageMUs.RLock()
	value := hashStorage[key]
	hashStorageMUs.RUnlock()
	return value, nil
}
