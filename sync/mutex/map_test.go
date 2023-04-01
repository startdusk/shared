package mutex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMutexMapLifecycle(t *testing.T) {
	m := NewMap[int, int]()

	// add a key value
	m.Set(1, 1)
	v, ok := m.Get(1)
	assert.Equal(t, true, ok)
	assert.Equal(t, v, 1)

	// update a key value
	m.Set(1, 2)
	v1, ok := m.Get(1)
	assert.Equal(t, true, ok)
	assert.Equal(t, v1, 2)

	length := m.Len()
	assert.Equal(t, 1, length)

	// delete a key value
	m.Delete(1)
	v2, ok := m.Get(1)
	assert.Equal(t, false, ok)
	assert.Equal(t, v2, 0)

	length1 := m.Len()
	assert.Equal(t, 0, length1)

	// range key value

	// clear map
}

func FuzzMutexMap(f *testing.F) {
	// TODO: test MutexMap
}
