package xctx_test

import (
	"Linda/baselibs/abstractions/xctx"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errTest = errors.New("test error")

func TestErrorHandler_Run(t *testing.T) {
	eh := xctx.NewErrHandler()
	eh.Run(
		func() {
			panic(errTest)
		})
	assert.Equal(t, errTest, eh.Err)
}
