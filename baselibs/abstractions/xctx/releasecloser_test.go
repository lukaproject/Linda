package xctx_test

import (
	"Linda/baselibs/abstractions/xctx"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCloser struct {
	isClosed bool
	Count    int
}

func (tc *testCloser) Close() error {
	tc.isClosed = true
	return nil
}

func (tc *testCloser) AddCount() {
	if !tc.isClosed {
		tc.Count++
	}
}

func TestReleaseCloser_Run(t *testing.T) {
	tc := testCloser{}
	xctx.NewCloser(&tc).Run(func() {
		tc.AddCount()
	})
	assert.Equal(t, 1, tc.Count)
	assert.True(t, tc.isClosed)
}
