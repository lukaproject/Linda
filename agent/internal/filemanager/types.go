package filemanager

type DownloadSourceType string

const (
	// file store in public url
	DownloadSourceType_public = "public"
	// file store in agentcentral local
	DownloadSourceType_internal = "internal"
)

func (dst DownloadSourceType) IsPublic() bool {
	return dst == DownloadSourceType_public
}

func (dst DownloadSourceType) IsInternal() bool {
	return dst == DownloadSourceType_internal
}

type DownloadInput struct {
	Type DownloadSourceType

	// for public download, the source url of public file.
	SourceURL string

	// download file from internal. Use block, filename to download from internal service.
	Block    string
	FileName string

	TargetPath string
}
