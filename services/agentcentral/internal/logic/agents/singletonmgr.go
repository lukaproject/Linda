package agents

var mgrInstance Mgr

func GetMgrInstance() Mgr {
	return mgrInstance
}

func InitMgrInstance() {
	mgrInstance = NewMgr()
}
