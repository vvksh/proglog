package log

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStoreWriteRead(t *testing.T) {
	storeFilename := "/tmp/0.store"
	os.Remove(storeFilename)

	store, err := newStore(storeFilename)
	require.NoError(t, err)

	//Write back to back 3 records
	data := []byte("hello world")
	pos, err := store.Write(data)
	data2 := []byte("hello world2")
	pos2, err := store.Write(data2)
	data3 := []byte("hello world3")
	pos3, err := store.Write(data3)
	require.NoError(t, err)

	// read all written data
	readData, err := store.Read(pos)
	require.NoError(t, err)
	require.Equal(t, data, readData)
	readData2, err := store.Read(pos2)
	assert.NoError(t, err)
	assert.Equal(t, data2, readData2)
	readData3, err := store.Read(pos3)
	assert.NoError(t, err)
	assert.Equal(t, data3, readData3)
}
