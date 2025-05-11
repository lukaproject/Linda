package xctx

// CtxReleaser
type CtxReleaser interface {
	// 保证在执行完函数f之后，释放掉CtxReleaser
	// 所持有的资源, 可以被其保管的类型如下
	// > sync.Locker
	// > io.Closer
	Run(f func())
}
