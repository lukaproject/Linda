package handler

import (
	"Linda/agent/internal/filemanager"
	"Linda/baselibs/abstractions/xlog"
	"Linda/protocol/models"
	"fmt"
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

func listFileInfos(h *Handler, logger xlog.Logger, fileMgr filemanager.Mgr, files []models.FileListRequest) {
	logger.Infof("start to list files %v", files)
	wg := &sync.WaitGroup{}
	n := len(files)
	wg.Add(n)
	responses := make([]models.FileListResponse, n)

	for id, f := range files {
		go func(responses []models.FileListResponse, idx int, req models.FileListRequest) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					logger.Errorf("panic in listFileInfos goroutine: %v", r)
					responses[idx] = models.FileListResponse{
						OperationId: req.OperationId,
						Error:       fmt.Sprintf("internal error: %v", r),
						Files:       nil,
					}
				}
			}()

			logger.Infof("listing files in directory: %s", req.DirPath)
			content, err := fileMgr.ListFileInfos(req.DirPath)

			var errorMsg string
			if err != nil {
				errorMsg = err.Error()
			}

			responses[idx] = models.FileListResponse{
				OperationId: req.OperationId,
				Error:       errorMsg,
				Files:       content,
			}
		}(responses, id, f)
	}

	wg.Wait()

	// Store responses for heartbeat transmission
	h.responsesMutex.Lock()
	for i, resp := range responses {
		if resp.Error != "" {
			logger.Errorf("list files from path %s failed, err=%v", files[i].DirPath, resp.Error)
		} else {
			logger.Infof("successfully listed %d files from %s", len(resp.Files), files[i].DirPath)
		}
		h.fileListResponses = append(h.fileListResponses, resp)
	}
	h.responsesMutex.Unlock()
}

func getFiles(h *Handler, logger xlog.Logger, fileMgr filemanager.Mgr, files []models.FileGetRequest) {
	logger.Infof("start to get files %v", files)
	wg := &sync.WaitGroup{}
	n := len(files)
	wg.Add(n)
	responses := make([]models.FileGetResponse, n)

	for id, f := range files {
		go func(responses []models.FileGetResponse, idx int, req models.FileGetRequest) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					logger.Errorf("panic in getFiles goroutine: %v", r)
					responses[idx] = models.FileGetResponse{
						OperationId: req.OperationId,
						Error:       fmt.Sprintf("internal error: %v", r),
						Content:     nil,
					}
				}
			}()

			content, err := fileMgr.GetFile(req.FilePath)

			var errorMsg string
			if err != nil {
				errorMsg = err.Error()
			}

			responses[idx] = models.FileGetResponse{
				OperationId: req.OperationId,
				Content:     content,
				Error:       errorMsg,
			}
		}(responses, id, f)
	}

	wg.Wait()

	// Store responses for heartbeat transmission
	h.responsesMutex.Lock()
	for i, resp := range responses {
		if resp.Error != "" {
			logger.Errorf("get file from path %s failed, err=%v", files[i].FilePath, resp.Error)
		} else {
			logger.Infof("successfully got file %s (%d bytes)", files[i].FilePath, len(resp.Content.Content))
		}
		h.fileGetResponses = append(h.fileGetResponses, resp)
	}
	h.responsesMutex.Unlock()
}
