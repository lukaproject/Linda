package errno

import "errors"

var (
	ErrEmptyBag = errors.New("no active task in bag")
)
