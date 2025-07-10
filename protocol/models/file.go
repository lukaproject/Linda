package models

type FileInfo struct {
	Name    string
	Path    string
	Size    int64
	ModTime int64
	IsDir   bool
}

type FileContent struct {
	FileInfo FileInfo
	Content  []byte
}
