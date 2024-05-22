package pkg

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreCommand(t *testing.T) {
	key := `1`
	val := `10`
	exp := 0
	store := NewStore()
	store.Set(key, val, int64(exp))
	v, err := store.Get(key)
	assert.Nil(t, err)
	assert.Equal(t, v, val)
	err = store.Del(key)
	assert.Nil(t, err)
	v, err = store.Get(key)
	assert.NotNil(t, err)
	assert.Nil(t, v)
	assert.EqualErrorf(t, err, "key not found", "Expected error 'key not found', got '%v'", err)
}

func TestStoreLoadFile(t *testing.T) {
	tmpDir := t.TempDir()
	filename := "test_store.dat"
	key := `1`
	val := `10`
	exp := 0
	store := NewStore()
	store.Set(key, val, int64(exp))
	v, err := store.Get(key)
	assert.Nil(t, err)
	assert.Equal(t, v, val)
	err = store.SaveToFile(path.Join(tmpDir, filename))
	assert.Nil(t, err)
	store = NewStore()
	err = store.LoadFromFile(path.Join(tmpDir, filename))
	assert.Nil(t, err)
	v, err = store.Get(key)
	assert.Nil(t, err)
	assert.Equal(t, v, val)
}
