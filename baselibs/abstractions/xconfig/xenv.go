package xconfig

import (
	"Linda/baselibs/abstractions/defaultor"
	"Linda/baselibs/abstractions/xref"
	"os"
	"reflect"
	"strconv"

	"github.com/lukaproject/xerr"
)

const (
	xenvTagKey = "xenv"
)

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

func NewFromOSEnv[T any]() *T {
	x := defaultor.New[T]()
	xref.WalkValues(x, walkToSetFromOSEnv)
	return x
}
