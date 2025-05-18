package xconfig

type CommConf struct {
	Env string `xdefault:"debug" xenv:"env"`
}
