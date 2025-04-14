package errno

import "errors"

var (
	ErrEmptyBag                = errors.New("no active task in bag")
	ErrNodeIdExists            = errors.New("nodeId exists")
	ErrAgentNotFound           = errors.New("agent not found")
	ErrNodeBelongsToAnotherBag = errors.New("node belongs to another bag")
	ErrInvalidDownloadType     = errors.New("no such download type")
	ErrTargetPathIsEmpty       = errors.New("target path is empty")
	ErrSourceURLIsEmpty        = errors.New("source URL is empty")
	ErrFileNameIsEmpty         = errors.New("input fileName should not be empty")
)
