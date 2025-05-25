package xstructs

// DepStruct
type DepStruct interface {
	// AddDeps 表示当前instance需要在这些deps全部都初始化成功后
	// 才可以开始初始化
	AddDeps(deps ...DepStruct)

	Init() error

	// Name 必须保证不同的instance有不同的Name，否则会出现冲突
	Name() string

	// GetDeps 获取当前instance的所有依赖instance name
	GetDeps() []string
}
