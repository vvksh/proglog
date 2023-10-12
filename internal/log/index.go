package log

import "os"

type index struct {
	file *os.File
}

func newIndex(f *os.File) (*index, error) {
	return &index{
		file: f,
	}, nil
}
