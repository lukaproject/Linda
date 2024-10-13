package filemanager

type Mgr interface {
	Download(url, targetPath string) error
	Exists(path string) bool
	Remove(path string) bool
	ListFiles(filesCh chan string) error
	Initial()
}
