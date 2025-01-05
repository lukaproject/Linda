package agents

import "Linda/baselibs/abstractions/xlog"

var (
	mgrInstance Mgr
	logger      = xlog.NewForPackage()
)

func GetMgrInstance() Mgr {
	return mgrInstance
}

func InitMgrInstance() {
	mgrInstance = NewMgr()
}
