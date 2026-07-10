package services

import "sync"

var (
	idempotencyStore = make(map[string]bool)
	mutex            sync.Mutex
)

func Exists(key string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	return idempotencyStore[key]
}

func Save(key string) {
	mutex.Lock()
	defer mutex.Unlock()

	idempotencyStore[key] = true
}