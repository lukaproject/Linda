package xctx

import "github.com/lukaproject/xerr"

type ErrorHandler struct {
	Err         error
	FinallyFunc func(err error)
}

func (eh *ErrorHandler) Run(f func()) {
	defer xerr.Recover(&eh.Err)
	if eh.FinallyFunc != nil {
		defer eh.FinallyFunc(eh.Err)
	}
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
