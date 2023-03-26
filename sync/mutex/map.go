package mutex

import (
	"sync"
)

// Map defines a mutex map
type Map[K comparable, V any] struct {
	mu sync.Mutex
	m  map[K]V
}

// NewMap creates a mutex map
func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		m: make(map[K]V),
	}
}

// Get gets a value
func (mm *Map[K, V]) Get(k K) (V, bool) {
	mm.mu.Lock()
	defer mm.mu.Unlock()
	val, ok := mm.m[k]
	return val, ok
}

// Set adds and updates key value
func (mm *Map[K, V]) Set(k K, v V) {
	mm.mu.Lock()
	defer mm.mu.Unlock()
	mm.m[k] = v
}

// Delete deletes a key value
func (mm *Map[K, V]) Delete(k K) {
	mm.mu.Lock()
	defer mm.mu.Unlock()
	delete(mm.m, k)
}

// Inner gets a inner map
func (mm *Map[K, V]) Inner() map[K]V {
	mm.mu.Lock()
	defer mm.mu.Unlock()
	return mm.inner()
}

// Range ranges a map with mutable ref
func (mm *Map[K, V]) Range(f func(store map[K]V, k K, v V)) {
	mm.mu.Lock()
	defer mm.mu.Unlock()
	innerMap := mm.inner()
	for k, v := range innerMap {
		f(mm.m, k, v)
	}
}

// Cap gets a map capacity
func (mm *Map[K, V]) Cap() int {
	mm.mu.Lock()
	defer mm.mu.Unlock()
	return len(mm.m)
}

func (mm *Map[K, V]) inner() map[K]V {
	newMap := make(map[K]V)
	for k, v := range mm.m {
		newMap[k] = v
	}
	return newMap
}
