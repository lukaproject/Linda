package apis

type AddBagReq struct {
	BagDisplayName string `example:"test-bagDisplayName"`
}

type AddTaskReq struct {
	TaskDisplayName string `example:"test-taskDisplayName"`
	ScriptPath      string `example:"/bin/test.sh"`
	WorkingDir      string `example:"/bin/testWorkingDir/working"`
}
