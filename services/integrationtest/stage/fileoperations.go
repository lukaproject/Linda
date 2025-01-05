package stage

import (
	"Linda/baselibs/apiscall/swagger"
	"Linda/baselibs/testcommon/gen"
	"context"
	"os"
	"path"
	"testing"

	"github.com/lukaproject/xerr"
)

type FileOperations struct {
	t   *testing.T
	cli *swagger.APIClient
}

func (fo *FileOperations) UploadFileContent(fileContent string, blockName string, targetFileName string) (err error) {
	tmpFilePath := path.Join(fo.t.TempDir(), xerr.Must(gen.StrGenerate(gen.CharsetLowerCase, 3, 10)))
	f := xerr.Must(os.Create(tmpFilePath))
	f.WriteString(fileContent)
	f.Close()
	_, _, err = fo.cli.FilesApi.FilesUploadPost(
		context.Background(),
		targetFileName,
		blockName,
		xerr.Must(os.Open(tmpFilePath)))
	return
}
