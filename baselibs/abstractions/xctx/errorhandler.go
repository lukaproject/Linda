package xctx

import "github.com/lukaproject/xerr"

type ErrorHandler struct {
	Err error
	// FinallyFunc 永远会在Recover之前执行
	FinallyFunc func()
}

func (eh *ErrorHandler) Run(f func()) {
	defer xerr.Recover(&eh.Err)
	if eh.FinallyFunc != nil {
		defer eh.FinallyFunc()
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
