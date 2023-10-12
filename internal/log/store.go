package log

import (
	"errors"
	"os"
	"sync"
)

type store struct {
	file *os.File
	size int64
	mu   sync.Mutex
}

func newStore(f *os.File) (*store, error) {
	fInfo, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}
	if fInfo.IsDir() {
		return nil, errors.New("got a directory")
	}
	size := fInfo.Size()
	return &store{
		file: f,
		size: size,
	}, nil
}

// Append takes byte array, appends to file and if write successful, returns num_bytes_written, pos ??, error
func (s *store) Append(p []byte) (int64, int64, error) {
	//TODO
	return 0,0 , nil
}
