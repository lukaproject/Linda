package fakefileserver_test

import (
	"Linda/baselibs/testcommon/fake/fakefileserver"
	"io"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/assert"
)

func TestFileServerAPI(t *testing.T) {
	s := fakefileserver.StartT(t)
	_ = s.AddFileContent(filepath.Join("test1", "ok", "test.txt"), "testcontent")
	resp := xerr.Must(
		http.Get(
			fakefileserver.BuildDownloadURL(
				s, filepath.Join("test1", "ok", "test.txt"))))
	result := string(xerr.Must(io.ReadAll(resp.Body)))
	assert.Equal(t, "testcontent", result)
}
