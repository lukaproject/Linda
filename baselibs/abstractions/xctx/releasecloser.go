package xctx

import "io"

type ReleaseCloser struct {
	closer io.Closer
	err    error
}

func (rc *ReleaseCloser) Run(f func()) {
	defer func() {
		rc.err = rc.closer.Close()
	}()
	f()
}

func NewCloser(closer io.Closer) *ReleaseCloser {
	return &ReleaseCloser{
		closer: closer,
	}
}
