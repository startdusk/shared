package mutex

import (
	"sync"

	"golang.org/x/exp/constraints"
)

type MutexMap[K constraints.Ordered, V any] struct {
	mu sync.Mutex
	m  map[K]V
}

func NewMutexMap[K constraints.Ordered, V any]() *MutexMap[K, V] {
	return &MutexMap[K, V]{
		m: make(map[K]V),
	}
}

func (mm *MutexMap[K, V]) Get(k K) (V, bool) {
	mm.mu.Lock()
	defer mm.mu.Unlock()
	val, ok := mm.m[k]
	return val, ok
}

func (mm *MutexMap[K, V]) Set(k K, v V) {
	mm.mu.Lock()
	defer mm.mu.Unlock()
	mm.m[k] = v
}

func (mm *MutexMap[K, V]) Delete(k K) {
	mm.mu.Lock()
	defer mm.mu.Unlock()
	delete(mm.m, k)
}

func (mm *MutexMap[K, V]) Inner() map[K]V {
	mm.mu.Lock()
	defer mm.mu.Unlock()
	return mm.inner()
}

func (mm *MutexMap[K, V]) Range(f func(store map[K]V, k K, v V)) {
	mm.mu.Lock()
	defer mm.mu.Unlock()
	innerMap := mm.inner()
	for k, v := range innerMap {
		f(mm.m, k, v)
	}
}

func (mm *MutexMap[K, V]) inner() map[K]V {
	newMap := make(map[K]V)
	for k, v := range mm.m {
		newMap[k] = v
	}
	return newMap
}
