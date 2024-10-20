package filemanager

type Mgr interface {
	Download(url, targetPath string) error
	Exists(path string) bool
	Remove(path string) bool
	ListFiles(filesCh chan string) error
	Initial()
}

type mgr struct {
	fileDownloader
}

func (fmgr *mgr) Download(url string, targetPath string) (err error) {
	err = fmgr.fileDownloader.Download(url, targetPath)
	return
}

func (fmgr *mgr) Exists(path string) bool {
	panic("not implemented") // TODO: Implement
}

func (fmgr *mgr) Remove(path string) bool {
	panic("not implemented") // TODO: Implement
}

func (fmgr *mgr) ListFiles(filesCh chan string) error {
	panic("not implemented") // TODO: Implement
}

func (fmgr *mgr) Initial() {
	panic("not implemented") // TODO: Implement
}

var (
	mgrInstance Mgr = nil
)

func Initial() {
	mgrInstance = &mgr{}
}

func MgrInstance() Mgr {
	return mgrInstance
}
