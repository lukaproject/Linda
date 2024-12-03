package xio_test

import (
	"Linda/baselibs/abstractions/xio"
	"io"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/suite"
)

type testReader struct {
	Offset   int
	Source   []byte
	BuffSize int
}

func (r *testReader) Read(p []byte) (n int, err error) {
	n = min(len(p), r.BuffSize, len(r.Source)-r.Offset)
	copy(p, r.Source[r.Offset:r.Offset+n])
	r.Offset += n
	if r.Offset == len(r.Source) {
		err = io.EOF
	}
	return
}

func (r *testReader) GenerateSource(n int) {
	r.Source = make([]byte, n)
	for i := 0; i < n; i++ {
		r.Source[i] = byte(rand.Intn(1 << 8))
	}
}

type testWriter struct {
	Dst []byte
}

func (w *testWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	w.Dst = append(w.Dst, p...)
	return
}

type transportTestSuite struct {
	suite.Suite
}

func (s *transportTestSuite) TestTransportSuccess() {
	tr := &testReader{
		Offset:   0,
		BuffSize: 1 << 10,
	}
	tr.GenerateSource(1 << 20)
	tw := &testWriter{
		Dst: make([]byte, 0),
	}
	s.Nil(xio.Transport(tr, tw))

	s.Equal(len(tr.Source), len(tw.Dst))
	for i := 0; i < len(tr.Source); i++ {
		s.Equal(tr.Source[i], tw.Dst[i])
	}
}

func TestTransportMain(t *testing.T) {
	suite.Run(t, new(transportTestSuite))
}
