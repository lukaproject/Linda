package tasks

var (
	bagsMgrInstance BagsMgr
)

func GetBagsMgrInstance() BagsMgr {
	return bagsMgrInstance
}

func InitBagsMgrInstance() {
	bagsMgrInstance = &bagsManager{}
}
