package utils

import (
	"errors"
	"runtime"
)

func GetDefaultShell() string {
	var shell string
	switch runtime.GOOS {
	case "windows":
		shell = "pwsh"
	case "linux":
		shell = "/bin/bash"
	default:
		panic(errors.New("unsupported OS"))
	}
	return shell
}
