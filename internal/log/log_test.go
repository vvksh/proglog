package log

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	api "github.com/vvksh/proglog/api/v1"
	"google.golang.org/protobuf/proto"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type LogTestSuite struct {
	suite.Suite
	log    *Log
	config Config
	dir    string
}

func (s *LogTestSuite) SetupSuite() {
	s.config = Config{}
	s.config.Segment.MaxStoreBytes = 32

}

// this function executes before each test case
func (s *LogTestSuite) SetupTest() {
	dir, err := ioutil.TempDir("", "store-test")
	require.NoError(s.T(), err)
	log, err := NewLog(dir, s.config)
	s.log = log
	s.Require().NoError(err)
}

// this function executes after each test case
func (s *LogTestSuite) TearDownTest() {
	os.RemoveAll(s.dir)
}

func TestLog(t *testing.T) {
	suite.Run(t, new(LogTestSuite))
}

func (s *LogTestSuite) TestAppendRead() {
	record := &api.Record{
		Value: []byte("hello world"),
	}
	offset, err := s.log.Append(record)
	s.Require().NoError(err)
	s.Require().Equal(uint64(0), offset)

	read, err := s.log.Read(offset)
	s.Require().NoError(err)
	s.Require().Equal(record.Value, read.Value)
}

func (s *LogTestSuite) testOutofRangeErr() {
	r, err := s.log.Read(1)
	s.Require().Nil(r)
	s.Require().Error(err)
}

// testInitExisting tests that we can reinstantiates log from existing segment files.
func (s *LogTestSuite) testInitExisting() {
	r := &api.Record{
		Value: []byte("hello world"),
	}
	for i := 0; i < 3; i++ {
		_, err := s.log.Append(r)
		s.Require().NoError(err)
	}
	s.Require().NoError(s.log.Close())
	off, err := s.log.LowestOffset()
	s.Require().NoError(err)
	s.Require().Equal(uint64(0), off)

	off, err = s.log.HighestOffset()
	s.Require().NoError(err)
	s.Require().Equal(uint64(2), off)

	s.log, err = NewLog(s.dir, s.config)
	s.Require().NoError(err)
	off, err = s.log.LowestOffset()
	s.Require().NoError(err)
	s.Require().Equal(uint64(0), off)

	off, err = s.log.HighestOffset()
	s.Require().NoError(err)
	s.Require().Equal(uint64(2), off)
}

func (s *LogTestSuite) testReader() {
	r := &api.Record{
		Value: []byte("hello world"),
	}
	off, err := s.log.Append(r)
	s.Require().NoError(err)
	s.Require().Equal(uint64(0), off)

	reader := s.log.Reader()
	b, err := ioutil.ReadAll(reader)
	s.Require().NoError(err)

	read := &api.Record{}
	err = proto.Unmarshal(b[lenWidth:], read)
	s.Require().NoError(err)
	s.Require().Equal(r.Value, read.Value)
}

func (s *LogTestSuite) testTruncate() {
	r := &api.Record{
		Value: []byte("hello world"),
	}
	for i := 0; i < 3; i++ {
		_, err := s.log.Append(r)
		s.Require().NoError(err)
	}

	err := s.log.Truncate(1)
	s.Require().NoError(err)

	_, err = s.log.Read(0)
	s.Require().Error(err)
}
