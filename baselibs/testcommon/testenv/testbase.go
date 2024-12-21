package testenv

import (
	"Linda/baselibs/abstractions/xos"
	"io/fs"
	"path"
	"strings"

	"github.com/stretchr/testify/suite"
)

// TestBase
type TestBase struct {
	suite.Suite

	// the base tmpdir for whole test suite
	tmpdir string
}

func (tb *TestBase) SetupSuite() {
	tb.tmpdir = tb.T().TempDir()
}

// TempDir
// Return the tempdir path of **subtest**.
func (tb *TestBase) TempDir() (subTestTmpDir string) {
	subTestFuncList := strings.Split(tb.T().Name(), "/")
	subTestTempDirName := subTestFuncList[len(subTestFuncList)-1]
	subTestTmpDir = path.Join(tb.tmpdir, subTestTempDirName)
	if !xos.PathExists(subTestTmpDir) {
		xos.MkdirAll(subTestTmpDir, fs.ModePerm)
	}
	if !xos.IsDir(subTestTmpDir) {
		tb.T().Fatalf("path %s is not a directory, pls check", subTestTmpDir)
	}
	return
}
