package filemanager

import (
	"Linda/baselibs/codes/errno"
	"Linda/protocol/models"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
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

	// ListFileInfos
	// will list all files and directories info in this dir.
	// only solved for absolute dir name.
	ListFileInfos(dirname string) ([]models.FileInfo, error)

	// GetFile
	// get file content and info by file path
	GetFile(filePath string) (*models.FileContent, error)

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
		panic("not implementation")
	} else {
		return errno.ErrInvalidDownloadType
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

// ListFileInfos calls the operator's ListFileInfos method
func (fmgr *mgr) ListFileInfos(dirname string) ([]models.FileInfo, error) {
	var fileInfos []models.FileInfo

	entries, err := os.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			logger.Errorf("failed to get info for %s: %v", entry.Name(), err)
			continue
		}

		fileInfo := models.FileInfo{
			Name:    entry.Name(),
			Path:    filepath.Join(dirname, entry.Name()),
			Size:    info.Size(),
			ModTime: info.ModTime().Unix(),
			IsDir:   entry.IsDir(),
		}

		fileInfos = append(fileInfos, fileInfo)
	}

	return fileInfos, nil
}

// GetFile calls the operator's GetFile method
func (fmgr *mgr) GetFile(filePath string) (*models.FileContent, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file not found: %s", filePath)
		}
		return nil, fmt.Errorf("failed to get file info: %v", err)
	}

	// Check if it's a directory
	if info.IsDir() {
		return nil, fmt.Errorf("path is a directory, not a file: %s", filePath)
	}

	// Read file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %v", err)
	}

	fileInfo := models.FileInfo{
		Name:    info.Name(),
		Path:    filePath,
		Size:    info.Size(),
		ModTime: info.ModTime().Unix(),
		IsDir:   info.IsDir(),
	}

	return &models.FileContent{
		FileInfo: fileInfo,
		Content:  content,
	}, nil
}

func (fmgr *mgr) Initial() {
	// NOTHING
}

func (fmgr *mgr) validateDownloadInput(input *DownloadInput) (err error) {
	if input.TargetPath == "" {
		return errno.ErrTargetPathIsEmpty
	}
	if input.Type.IsPublic() {
		if input.SourceURL == "" {
			return errno.ErrSourceURLIsEmpty
		}
	} else if input.Type.IsInternal() {
		if input.Block == "" {
			return errors.New("input block should not be empty")
		}
		if input.FileName == "" {
			return errno.ErrFileNameIsEmpty
		}
	}
	return
}

func NewMgr() Mgr {
	fmgr := &mgr{}
	fmgr.Initial()
	return fmgr
}
