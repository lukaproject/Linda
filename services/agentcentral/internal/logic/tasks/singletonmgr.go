package tasks

var mgrInstance Mgr

func GetMgrInstance() Mgr {
	return mgrInstance
}

func InitMgrInstance() {
	mgrInstance = &manager{}
}
