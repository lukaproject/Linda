package db

import "Linda/baselibs/abstractions/xlog"

var (
	logger = xlog.NewForPackage()
)

type CountType struct {
	Count uint32
}
