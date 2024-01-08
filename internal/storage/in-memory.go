package storage

import "sync"

var storage = make(map[string]string)
var storageMUs = sync.RWMutex{}

func Set(key string, value string) error {
	storageMUs.Lock()
	storage[key] = value
	storageMUs.Unlock()
	return nil
}

func Get(key string) (string, error) {
	storageMUs.RLock()
	value := storage[key]
	storageMUs.RUnlock()
	return value, nil
}
