package gen_test

import (
	"Linda/baselibs/abstractions/xos"
	"Linda/baselibs/testcommon/gen"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileGen(t *testing.T) {
	tmpdir := t.TempDir()
	xos.MkdirAll(tmpdir, fs.ModePerm)
	roles := gen.FileGenerateRoles{
		MaxDirDepth: 3,
		MaxNameLen:  5,
		MaxCount:    30,
		RootDir:     tmpdir,
	}
	assert.Nil(t, gen.FileGenerate(roles))

	curMaxNameLen := 0
	curMaxDepth := 0
	curCount := 0
	assert.Nil(t, filepath.Walk(tmpdir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		dirpath := path
		if !xos.IsDir(path) {
			dirpath = filepath.Dir(path)
			curCount++
		}
		relativePath, err := filepath.Rel(tmpdir, dirpath)
		if err != nil {
			return err
		}
		curMaxDepth = max(curMaxDepth, 1+strings.Count(relativePath, "/"))
		return nil
	}))
	assert.LessOrEqual(t, curCount, roles.MaxCount)
	assert.LessOrEqual(t, curMaxDepth, roles.MaxDirDepth)
	assert.LessOrEqual(t, curMaxNameLen, roles.MaxNameLen)
}
