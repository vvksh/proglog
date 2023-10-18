package log

/**
a log stores incoming records into a segment. A segment consists of a store file to store the actual records and an index file which has mapping record_id to offset in store file
once the segment reaches a certain size, it closes the segment and creates a new one for store the new records.

we create the log by passing a directory name where all current segment and old segment files live.

Functionalities:
- append a record -> append record to current segment
- read a record given id -> has to scan through all the segments and check each one
**/
import (
	"io"

	api "github.com/vvksh/proglog/api/v1"
)

type Log struct {
	Dir            string
	currentSegment *segment
	segments       []*segment
	Config
}

const (
	DEFAULT_MAX_STORE_BYTES = 1024
	DEFAULT_INDEX_BYTES     = 1024
	DEFAULT_INITIAL_OFFSET  = 0
)

func NewLog(dir string, c Config) (*Log, error) {
	if c.Segment.MaxStoreBytes == 0 {
		c.Segment.MaxStoreBytes = DEFAULT_MAX_STORE_BYTES
	}
	if c.Segment.MaxIndexBytes == 0 {
		c.Segment.MaxIndexBytes = DEFAULT_INDEX_BYTES
	}

	l := &Log{
		Dir:    dir,
		Config: c,
	}
	return l, l.setup()
}

func (l *Log) setup() error {
	// files := ioutil.ReadDir(l.Dir)
	// if
	//TODO read all  files in the dir and initialize segments 
	return nil

}

func (l *Log) Append(record *api.Record) (uint64, error) {
	return 0, nil
}

func (l *Log) Read(offset uint64) (*api.Record, error) {
	//TODO
	return &api.Record{}, nil
}

func (l *Log) Close() error {
	//TODO
	return nil
}

func (l *Log) LowestOffset() (uint64, error) {
	//TODO
	return uint64(0), nil
}

func (l *Log) HighestOffset() (uint64, error) {
	//TODO
	return uint64(0), nil
}

func (l *Log) Reader() io.Reader {
	//TODO
	return nil
}

func (l *Log) Truncate(off uint64) error {
	// TODO
	return nil
}