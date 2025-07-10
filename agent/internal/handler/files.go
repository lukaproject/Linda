package handler

import (
	"Linda/agent/internal/filemanager"
	"Linda/baselibs/abstractions/xlog"
	"Linda/protocol/models"
	"sync"
)

func downloadFiles(logger xlog.Logger, fileMgr filemanager.Mgr, files []models.FileDescription) {
	logger.Infof("start to download files %v", files)
	wg := &sync.WaitGroup{}
	n := len(files)
	wg.Add(n)
	errs := make([]error, n)

	for id, f := range files {
		go func(errs []error, idx int, file models.FileDescription) {
			defer wg.Done()
			errs[idx] = fileMgr.Download(filemanager.DownloadInput{
				Type:       filemanager.DownloadSourceType_public,
				SourceURL:  file.Uri,
				TargetPath: file.LocationPath,
			})
		}(errs, id, f)
	}

	for i, err := range errs {
		if err != nil {
			logger.Errorf("download %s failed, err=%v", files[i].Uri, err)
		}
	}
	wg.Wait()
}

func listFileInfos(logger xlog.Logger, fileMgr filemanager.Mgr, files []models.FileListRequest) {
	logger.Infof("start to download files %v", files)
	wg := &sync.WaitGroup{}
	n := len(files)
	wg.Add(n)
	responses := make([]models.FileListResponse, n)

	for id, f := range files {
		go func(responses []models.FileListResponse, idx int, req models.FileListRequest) {
			defer wg.Done()
			content, err := fileMgr.ListFileInfos(req.DirPath)
			responses[idx] = models.FileListResponse{
				OperationId: req.OperationId,
				Error:       err.Error(),
				Files:       content,
			}
		}(responses, id, f)
	}

	for i, resp := range responses {
		if resp.Error != "" {
			logger.Errorf("list files from path %s failed, err=%v", files[i].DirPath, resp.Error)
		}
	}
	wg.Wait()
}

func getFiles(logger xlog.Logger, fileMgr filemanager.Mgr, files []models.FileGetRequest) {
	logger.Infof("start to download files %v", files)
	wg := &sync.WaitGroup{}
	n := len(files)
	wg.Add(n)
	responses := make([]models.FileGetResponse, n)

	for id, f := range files {
		go func(responses []models.FileGetResponse, idx int, req models.FileGetRequest) {
			defer wg.Done()
			content, err := fileMgr.GetFile(req.FilePath)
			responses[idx] = models.FileGetResponse{
				OperationId: req.OperationId,
				Content:     content,
				Error:       err.Error(),
			}
		}(responses, id, f)
	}

	for i, resp := range responses {
		if resp.Error != "" {
			logger.Errorf("get file from path %s failed, err=%v", files[i].FilePath, resp.Error)
		}
	}
	wg.Wait()
}
