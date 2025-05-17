package xctx

import "github.com/lukaproject/xerr"

type ErrorHandler struct {
	Err error
}

func (eh *ErrorHandler) Run(f func()) {
	defer xerr.Recover(&eh.Err)
	f()
}

func NewErrHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func NewErrHandleRun(f func()) *ErrorHandler {
	eh := NewErrHandler()
	eh.Run(f)
	return eh
}
