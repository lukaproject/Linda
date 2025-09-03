package filemanager

import "Linda/protocol/models"

// MockMgr is a mock struct for filemanager.Mgr
type MockMgr struct {
	CallCount struct {
		Download      int
		Exists        int
		Remove        int
		ListFiles     int
		ListFileInfos int
		GetFile       int
		Initial       int
	}

	MockFuncs struct {
		Download      func(input DownloadInput) error
		Exists        func(path string) bool
		Remove        func(path string) error
		ListFiles     func(dirname string, filesCh chan string) error
		ListFileInfos func(dirname string) ([]models.FileInfo, error)
		GetFile       func(filePath string) (*models.FileContent, error)

		Initial func()
	}
}

// Download
// there are many source file types, if you want to download file from
// agentcentral, you can use internal source type, or you can provide a
// public reachable url to download. the TargetPath are required.
func (mkm *MockMgr) Download(input DownloadInput) error {
	mkm.CallCount.Download++
	if mkm.MockFuncs.Download != nil {
		return mkm.MockFuncs.Download(input)
	}
	return nil
}

// Exists
// Check if this path is exists.
func (mkm *MockMgr) Exists(path string) bool {
	mkm.CallCount.Exists++
	if mkm.MockFuncs.Exists != nil {
		return mkm.MockFuncs.Exists(path)
	}
	return false
}

// Remove
// delete all files in this path.
func (mkm *MockMgr) Remove(path string) error {
	mkm.CallCount.Remove++
	if mkm.MockFuncs.Remove != nil {
		return mkm.MockFuncs.Remove(path)
	}
	return nil
}

// ListFiles
// will list all files in this dir.
// only solved for absolute dir name.
func (mkm *MockMgr) ListFiles(dirname string, filesCh chan string) error {
	mkm.CallCount.ListFiles++
	if mkm.MockFuncs.ListFiles != nil {
		return mkm.MockFuncs.ListFiles(dirname, filesCh)
	}
	return nil
}

// ListFileInfos
func (mkm *MockMgr) ListFileInfos(dirname string) ([]models.FileInfo, error) {
	mkm.CallCount.ListFileInfos++
	if mkm.MockFuncs.ListFileInfos != nil {
		return mkm.MockFuncs.ListFileInfos(dirname)
	}
	return nil, nil
}

// GetFile
func (mkm *MockMgr) GetFile(filePath string) (*models.FileContent, error) {
	mkm.CallCount.GetFile++
	if mkm.MockFuncs.GetFile != nil {
		return mkm.MockFuncs.GetFile(filePath)
	}
	return nil, nil
}

func (mkm *MockMgr) Initial() {
}
