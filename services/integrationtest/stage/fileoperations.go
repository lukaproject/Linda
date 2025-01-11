package stage

import (
	"Linda/baselibs/apiscall/swagger"
	"Linda/baselibs/multifs"
	"Linda/baselibs/testcommon/gen"
	"os"
	"path"
	"testing"

	"github.com/lukaproject/xerr"
)

type FileOperations struct {
	t   *testing.T
	cli *swagger.APIClient

	fileServiceFEEndPoint string
}

func (fo *FileOperations) Upload(fileContent, targetFilePath string) (err error) {
	fileName := xerr.Must(gen.StrGenerate(gen.CharsetLowerCase, 3, 10))
	tmpFilePath := path.Join(fo.t.TempDir(), fileName)
	f := xerr.Must(os.Create(tmpFilePath))
	f.WriteString(fileContent)
	f.Close()

	f = xerr.Must(os.Open(tmpFilePath))
	defer f.Close()

	c := multifs.NewClient(fo.fileServiceFEEndPoint)
	if err = c.Upload(targetFilePath, fileName, f); err != nil {
		return
	}
	return
}
