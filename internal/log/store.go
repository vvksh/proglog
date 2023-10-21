package log

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"sync"
)

const (
	lenWidth = 8 // 64 bits

)

var enc = binary.BigEndian

type store struct {
	file    *os.File
	nextPos uint64
	mu      sync.Mutex
}

func newStore(filename string) (*store, error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	fInfo, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}
	if fInfo.IsDir() {
		return nil, errors.New("got a directory")
	}
	size := fInfo.Size()
	return &store{
		file:    f,
		nextPos: uint64(size),
	}, nil
}

// Write takes byte array, appends to file and if write successful, returns position of buffer and error
// First write the length of bytes in first 64 bits
func (s *store) Write(p []byte) (uint64, error) {
	buf := bytes.NewBuffer(make([]byte, 0, lenWidth))
	err := binary.Write(buf, enc, uint64(len(p)))
	if err != nil {
		return 0, err
	}
	fmt.Printf("% x", buf.Bytes())

	toWrite := append(buf.Bytes(), p...)
	s.mu.Lock()
	defer s.mu.Unlock()
	numBytesWritten, err := s.file.Write(toWrite)
	if err != nil {
		return 0, err
	}
	if err := s.file.Sync(); err != nil {
		return 0, err
	}
	pos := s.nextPos
	s.nextPos += uint64(numBytesWritten)
	return pos, nil
}

func (s *store) Read(pos uint64) ([]byte, error) {
	// read the length of bytes to read
	sizeBuf := make([]byte, lenWidth)
	_, err := s.file.ReadAt(sizeBuf, int64(pos))
	if err != nil {
		return []byte{}, err
	}
	sizeReader := bytes.NewReader(sizeBuf)
	var bytesLen uint64
	err = binary.Read(sizeReader, enc, &bytesLen)
	if err != nil {
		return []byte{}, err
	}

	// now read the bytes
	bytesBuf := make([]byte, bytesLen)
	dataPos := int64(pos) + lenWidth
	_, err = s.file.ReadAt(bytesBuf, dataPos)
	if err != nil {
		return []byte{}, err
	}
	return bytesBuf, nil
}
