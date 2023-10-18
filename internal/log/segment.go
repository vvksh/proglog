package log

import (
	api "github.com/vvksh/proglog/api/v1"
)

type segment struct {
	initialOffset uint64
	s             *store
	idx           *index
	Config
}

func newSegment(dir string, offset uint64, config Config) (*segment, error) {
	//TODO
	// - open a file and hold the file pointer for reading and appending
	// - open and initialize the index file
	return nil, nil
}

// Append adds a new record to the end of the store file and updates the corresponding index
// Its up to the caller to check for size limit, can use IsMaxed method
func (s *segment) Append(record *api.Record) error {
	// TODO
	return nil
}

func (s *segment) Read(offset uint64) (*api.Record, error) {
	// TODO
	return &api.Record{}, nil
}