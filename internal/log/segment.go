package log

import (
	"os"
	"path"
)

type segment struct {
	s   *store
	idx *index
}

func newSegment(dir string) (*segment, error) {
	//TODO
	// scan directory, pick next number to name the segment eg: 1, 2, 3 etc
	idx := "1"
	storeFile, err := os.Create(path.Join(dir, idx, ".store"))
	if err != nil {
		return nil, err
	}
	idxFile, err := os.Create(path.Join(dir, idx, ".idx"))
	if err != nil {
		return nil, err
	}
	store, err := newStore(storeFile)
	if err != nil {
		return nil, err
	}
	store, err := newStore(storeFile)
	if err != nil {
		return nil, err
	}

	return &segment{

	}
}
