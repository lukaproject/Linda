package xctx_test

import (
	"Linda/baselibs/abstractions/xctx"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errTest = errors.New("test error")

func TestErrorHandler_Run(t *testing.T) {
	eh := xctx.NewErrHandleRun(
		func() {
			panic(errTest)
		})
	assert.Equal(t, errTest, eh.Err)
}

func TestErrorHandler_RunWithFinallyFunc(t *testing.T) {
	count := 0
	eh := xctx.ErrorHandler{
		FinallyFunc: func() {
			count++
		},
	}

	eh.Run(func() {
		panic(errTest)
	})

	assert.Equal(t, errTest, eh.Err)
	assert.Equal(t, 1, count)
}
