package task

type MockMgr struct {
}

func (mm *MockMgr) AddTask(taskName string) {

}
func (mm *MockMgr) PopFinishedTasks() (finishedTaskNames []string) {
	return
}
