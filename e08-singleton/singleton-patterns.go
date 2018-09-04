package main

import (
	"sync"
	"sync/atomic"
)

var instance *int

func GetInstanceMistake() *int {
	if instance == nil {
		a := 0
		instance = &a // <--- NOT THREAD SAFE
	}
	return instance
}

var mu sync.Mutex

func GetInstanceSlow() *int {
	mu.Lock() // <--- Unnecessary locking if instance already created
	defer mu.Unlock()

	if instance == nil {
		a := 1
		instance = &a
	}
	return instance
}

func GetInstanceFlaw() *int {
	if instance == nil { // <-- Not yet perfect. since it's not fully atomic
		mu.Lock()
		defer mu.Unlock()

		if instance == nil {
			a := 1
			instance = &a
		}
	}
	return instance
}

var initialized uint32

func GetInstanceLong() *int {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		a := 1
		instance = &a
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}
