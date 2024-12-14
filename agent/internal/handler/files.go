package handler

import (
	"Linda/agent/internal/filemanager"
	"Linda/protocol/models"
	"sync"

	"github.com/sirupsen/logrus"
)

func downloadFiles(fileMgr filemanager.Mgr, files []models.FileDescription) {
	logrus.Debugf("start to download files %v", files)
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
			logrus.Errorf("download %s failed, err=%v", files[i].Uri, err)
		}
	}
	wg.Wait()
}
