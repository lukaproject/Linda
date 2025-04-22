package xconfig

import (
	"Linda/baselibs/abstractions/defaultor"
	"Linda/baselibs/abstractions/ds"
	"Linda/baselibs/abstractions/xref"
	"os"
	"reflect"
	"strconv"

	"github.com/lukaproject/xerr"
)

const (
	xenvTagKey = "xenv"
)

// NewFromOSEnv
// 读取环境变量中的值，设置在对应的xenv tag上
func NewFromOSEnv[T any]() *T {
	x := defaultor.New[T]()
	xref.WalkValues(x, walkToSetFromOSEnv)
	return x
}

// GetEnvs
// 获取这个type所有的环境变量，并返回
func GetEnvs[T any]() (envs ds.Set[string]) {
	envs = make(ds.Set[string])
	walkToGetOSEnv := func(input xref.WalkFuncInput) {
		_, tags := input.Value, input.FieldTag
		tagValue, ok := tags.Lookup(xenvTagKey)
		if !ok {
			return
		}
		envs.Insert(tagValue)
	}
	xref.WalkValues(defaultor.New[T](), walkToGetOSEnv)
	return
}

func walkToSetFromOSEnv(input xref.WalkFuncInput) {
	v, tags := input.Value, input.FieldTag

	tagValue, ok := tags.Lookup(xenvTagKey)
	if !ok {
		return
	}
	envValue, ok := os.LookupEnv(tagValue)
	if !ok {
		return
	}
	if v.CanSet() {
		switch v.Kind() {
		case reflect.String:
			v.SetString(envValue)
		case reflect.Float64, reflect.Float32:
			bitSize := 32
			if v.Kind() == reflect.Float64 {
				bitSize = 64
			}
			v.SetFloat(xerr.Must(strconv.ParseFloat(envValue, bitSize)))
		case reflect.Int:
			v.SetInt(xerr.Must(strconv.ParseInt(envValue, 10, 64)))
		case reflect.Bool:
			v.SetBool(xerr.Must(strconv.ParseBool(envValue)))
		}
	}
}
