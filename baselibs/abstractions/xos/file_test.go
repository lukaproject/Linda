package xos_test

import (
	"Linda/baselibs/abstractions/xos"
	"io/fs"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTouch(t *testing.T) {
	tmpdir := t.TempDir()
	testFile := path.Join(tmpdir, "testpath", "testfile")
	assert.False(t, xos.PathExists(testFile))
	assert.Nil(t, xos.Touch(testFile))
	assert.True(t, xos.PathExists(testFile))
}

func TestPathExists(t *testing.T) {
	tmpdir := t.TempDir()
	testFile := path.Join(tmpdir, "testpath", "testfile")
	assert.False(t, xos.PathExists(testFile))
	assert.True(t, xos.PathExists(tmpdir))
}

func TestIsDir(t *testing.T) {
	tmpdir := t.TempDir()
	testDir := path.Join(tmpdir, "testpath", "testdir")
	assert.False(t, xos.PathExists(testDir))
	xos.MkdirAll(testDir, fs.ModePerm)
	assert.True(t, xos.PathExists(testDir))
}
