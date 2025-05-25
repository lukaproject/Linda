package xstructs

// DepStructHelper 这是一个用来描述depStruct之间依赖关系的类型
// 使其可以不用独立实现AddDeps, Name和GetDeps这些函数，只需要组合这个helper即可
type DepStructHelper struct {
	NameStr string
	Deps    []string
}

func (dsh *DepStructHelper) AddDeps(deps ...DepStruct) {
	if dsh.Deps == nil {
		dsh.Deps = make([]string, 0)
	}
	for _, dep := range deps {
		dsh.Deps = append(dsh.Deps, dep.Name())
	}
}

func (dsh *DepStructHelper) GetDeps() []string {
	return dsh.Deps
}

func (dsh *DepStructHelper) Name() string {
	return dsh.NameStr
}
