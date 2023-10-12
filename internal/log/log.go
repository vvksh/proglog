package log

/**
a log stores incoming records into a segment. A segment consists of a store file to store the actual records and an index file which has mapping record_id to offset in store file
once the segment reaches a certain size, it closes the segment and creates a new one for store the new records.

we create the log by passing a directory name where all current segment and old segment files live
**/

type Log struct {
	s *segment
}


func NewLog(dir string) (*Log, error) {
	// instantiate a new segment
	segment, err  := newSegment(dir)
	if err != nil {
		return nil, err
	}
	return &Log{
		s: segment,
	}

}

