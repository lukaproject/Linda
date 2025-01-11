package multifs_test

import (
	"Linda/baselibs/multifs"
	"Linda/baselibs/testcommon/gen"
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/assert"
)

func TestUpload(t *testing.T) {
	localFileService := multifs.New(
		multifs.NewFileServiceInput{
			Port:    5555,
			BaseDir: t.TempDir(),
			Type:    multifs.FileServiceType_Local,
		})
	localFileService.Start()
	defer localFileService.Shutdown(context.Background())

	c := multifs.NewClient("http://localhost:5555")
	testContent := "test"
	reader := bytes.NewBuffer([]byte(testContent))
	err := c.Upload("test/test1/test.txt", "test.txt", reader)
	assert.Nil(t, err)

	getFile := bytes.NewBuffer([]byte{})
	err = c.DownloadToWriter("test/test1/test.txt", getFile)
	assert.Nil(t, err)
	fileContent := string(xerr.Must(io.ReadAll(getFile)))
	assert.Equal(t, testContent, fileContent)
	t.Log(fileContent)
}

func TestUploadLargeFile(t *testing.T) {
	localFileService := multifs.New(
		multifs.NewFileServiceInput{
			Port:    5555,
			BaseDir: t.TempDir(),
			Type:    multifs.FileServiceType_Local,
		})
	localFileService.Start()
	defer localFileService.Shutdown(context.Background())

	c := multifs.NewClient("http://localhost:5555")
	charset := gen.CharsetDigit + gen.CharsetLowerCase + gen.CharsetUpperCase
	testContent, err := gen.StrGenerate(charset, 1<<15, 1<<20)
	assert.Nil(t, err)
	reader := bytes.NewBuffer([]byte(testContent))
	err = c.Upload("test/test1/test.txt", "test.txt", reader)
	assert.Nil(t, err)

	getFile := bytes.NewBuffer([]byte{})
	err = c.DownloadToWriter("test/test1/test.txt", getFile)
	assert.Nil(t, err)
	fileContent := string(xerr.Must(io.ReadAll(getFile)))
	assert.Equal(t, testContent, fileContent)
	t.Log(fileContent[:50] + "...")
}
