package xctx_test

import (
	"Linda/baselibs/abstractions/xctx"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReleaseLocker_Run(t *testing.T) {
	mut := sync.Mutex{}
	value := 0
	xctx.NewLocker(&mut).Run(func() {
		value++
	})
	assert.True(t, mut.TryLock())
}
