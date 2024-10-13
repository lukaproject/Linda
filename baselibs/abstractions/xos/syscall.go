package xos

import (
	"os"
	"path"

	"github.com/lukaproject/xerr"
)

// CurrentPath
// 获取当前可执行文件所在的文件夹
func CurrentPath() string {
	return path.Dir(xerr.Must(os.Executable()))
}
