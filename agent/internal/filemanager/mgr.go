package filemanager

import (
	"errors"
	"os"
	"path"
)

// Mgr
// This is a filemanager Mgr manage Linda agent related files.
// it will related to each files it can touched.
type Mgr interface {
	// Download
	// there are many source file types, if you want to download file from
	// agentcentral, you can use internal source type, or you can provide a
	// public reachable url to download. the TargetPath are required.
	Download(input DownloadInput) error

	// Exists
	// Check if this path is exists.
	Exists(path string) bool

	// Remove
	// delete all files in this path.
	Remove(path string) error

	// ListFiles
	// will list all files in this dir.
	// only solved for absolute dir name.
	ListFiles(dirname string, filesCh chan string) error
	Initial()
}

type mgr struct {
	Downloader
	Operator
}

func (fmgr *mgr) Download(input DownloadInput) (err error) {
	if err = fmgr.validateDownloadInput(&input); err != nil {
		return
	}
	if input.Type.IsPublic() {
		return fmgr.DownloadFromPublicURL(input.SourceURL, input.TargetPath)
	} else if input.Type.IsInternal() {
		return fmgr.DownloadFromInternal(input.Block, input.FileName, input.TargetPath)
	} else {
		return errors.New("no such download type")
	}
}

func (fmgr *mgr) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (fmgr *mgr) Remove(path string) error {
	return os.RemoveAll(path)
}

func (fmgr *mgr) ListFiles(dir string, filesCh chan string) error {
	if !path.IsAbs(dir) {
		return errors.New("not absolute dir name")
	}
	fmgr.ListFileNames(dir, filesCh)
	return nil
}

func (fmgr *mgr) Initial() {
	// NOTHING
}

func (fmgr *mgr) validateDownloadInput(input *DownloadInput) (err error) {
	if input.TargetPath == "" {
		return errors.New("target path is empty")
	}
	if input.Type.IsPublic() {
		if input.SourceURL == "" {
			return errors.New("source URL is empty")
		}
	} else if input.Type.IsInternal() {
		if input.Block == "" {
			return errors.New("input block should not be empty")
		}
		if input.FileName == "" {
			return errors.New("input fileName should not be empty")
		}
	}
	return
}

func NewMgr() Mgr {
	fmgr := &mgr{}
	fmgr.Initial()
	return fmgr
}
