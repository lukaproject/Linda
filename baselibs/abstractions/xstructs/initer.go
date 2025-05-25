package xstructs

type Initer interface {
	// Init
	// call it before you use this instance.
	Init() error

	// Name
	// 这个函数是用来获取struct的名字
	// 用于唯一指定一个struct
	Name() string
}
