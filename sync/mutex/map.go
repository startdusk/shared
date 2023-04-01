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
func (m *Map[K, V]) Get(k K) (V, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	val, ok := m.m[k]
	return val, ok
}

// Set adds and updates key value
func (m *Map[K, V]) Set(k K, v V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[k] = v
}

// Delete deletes a key value
func (m *Map[K, V]) Delete(k K) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.m, k)
}

// Inner gets a inner map
func (m *Map[K, V]) Inner() map[K]V {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.inner()
}

// Range ranges a map with mutable ref
func (m *Map[K, V]) Range(f func(store map[K]V, k K, v V)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	innerMap := m.inner()
	for k, v := range innerMap {
		f(m.m, k, v)
	}
}

// Len gets a map length
func (m *Map[K, V]) Len() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.m)
}

// Swap replace the inner map
func (m *Map[K, V]) Swap(newMap map[K]V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if newMap == nil {
		m.Clear()
	} else {
		m.m = newMap
	}
}

// Clear renew the map
func (m *Map[K, V]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m = make(map[K]V)
}

func (m *Map[K, V]) inner() map[K]V {
	newMap := make(map[K]V)
	for k, v := range m.m {
		newMap[k] = v
	}
	return newMap
}
