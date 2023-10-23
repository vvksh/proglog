package log

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	api "github.com/vvksh/proglog/api/v1"
)

func TestSegment_Append_And_Read(t *testing.T) {
	dir, _ := os.MkdirTemp("", "segment-test")
	defer os.RemoveAll(dir)
	r := &api.Record{
		Value: []byte("hello world"),
	}

	c := Config{}
	c.Segment.MaxStoreBytes = 1024
	c.Segment.MaxIndexBytes = entWidth * 3

	s, err := newSegment(dir, 16, c)
	require.NoError(t, err)

	require.Equal(t, uint64(16), s.nextOffset)
	require.False(t, s.IsMaxed())

	for i := 0; i < 3; i++ {
		off, err := s.Append(r)
		require.NoError(t, err)
		require.Equal(t, 16+i, off)

		got, err := s.Read(off)
		require.NoError(t, err)
		require.Equal(t, r.Value, got.Value)
	}

	// should max out after 3 entries
	_, err = s.Append(r)
	require.Equal(t, io.EOF, err)
	require.True(t, s.IsMaxed())

	// try maxing out store file
	c.Segment.MaxStoreBytes = uint64(len(r.Value) * 3)
	c.Segment.MaxIndexBytes = 1024

	s, err = newSegment(dir, 16, c)
	require.NoError(t, err)

	// the segment should have restored with the existing dir, and with new config should have maxed out the store file
	// appending new record should result in a EOF error
	require.True(t, s.IsMaxed())

	err = s.Remove()
	require.NoError(t, err)
	s, err = newSegment(dir, 16, c)
	require.NoError(t, err)
	require.False(t, s.IsMaxed())
}
