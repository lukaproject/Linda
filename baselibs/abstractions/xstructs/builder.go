package xstructs

import "errors"

// Builder 是一个用来初始化所有DepStruct的容器
// 它可以保证所有的DepStruct按照依赖顺序进行启动
type Builder struct {
	depStructsMap map[string]DepStruct
}

// AddDepStructs 将 depStructs 加入当前容器进行管理
// Name相同的无法重复加入,会以第一次使用当前Name的struct为准
func (ds *Builder) AddDepStructs(depStructs ...DepStruct) {
	for _, depStruct := range depStructs {
		if _, ok := ds.depStructsMap[depStruct.Name()]; ok {
			continue
		}
		ds.depStructsMap[depStruct.Name()] = depStruct
	}
}

func (ds *Builder) Build() error {
	que := make([]string, 0)
	topo := make(map[string]int, len(ds.depStructsMap))
	graph := make(map[string][]string)

	for k, v := range ds.depStructsMap {
		deps := v.GetDeps()
		depsCount := len(deps)
		topo[k] = depsCount
		if depsCount == 0 {
			que = append(que, k)
		}
		for _, src := range deps {
			edges, ok := graph[src]
			if ok {
				graph[src] = append(edges, k)
			} else {
				edges = make([]string, 1)
				edges[0] = k
				graph[src] = edges
			}
		}
	}
	initCount := 0
	for len(que) > 0 {
		cur := que[0]
		ds.depStructsMap[cur].Init()
		initCount++
		que = que[1:]
		for _, target := range graph[cur] {
			topo[target]--
			if topo[target] <= 0 {
				que = append(que, target)
			}
		}
	}

	if initCount == len(ds.depStructsMap) {
		return nil
	}

	return errors.New("cycling dependencies")
}

// NewDepStructsBuilder 新建一个DepStructs 的 Builder
func NewDepStructsBuilder() *Builder {
	return &Builder{
		depStructsMap: make(map[string]DepStruct),
	}
}
