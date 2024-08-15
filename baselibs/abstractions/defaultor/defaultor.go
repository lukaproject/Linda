package defaultor

import (
	"Linda/baselibs/abstractions/xref"
	"reflect"
	"strconv"
	"strings"

	"github.com/lukaproject/xerr"
)

var (
	defaultorTagKey = "xdefault"
)

func toIntSlice(strs []string) (arr []int) {
	arr = make([]int, 0, len(strs))
	for _, v := range strs {
		arr = append(arr, int(xerr.Must(strconv.ParseInt(v, 10, 64))))
	}
	return
}

func parseStrsToSlice(t reflect.Type, v reflect.Value, strs []string) {
	switch t.Elem().Kind() {
	case reflect.String:
		v.Set(reflect.ValueOf(strs))
	case reflect.Int:
		v.Set(reflect.ValueOf(toIntSlice(strs)))
	case reflect.Float32, reflect.Float64:
	default:
	}
}

func walkToSetDefaultValues(fieldName string, tags reflect.StructTag, t reflect.Type, v reflect.Value) {
	tagValue, ok := tags.Lookup(defaultorTagKey)
	if !ok {
		return
	}
	if v.CanSet() {
		switch v.Kind() {
		case reflect.String:
			v.SetString(tagValue)
		case reflect.Float32, reflect.Float64:
			v.SetFloat(xerr.Must(strconv.ParseFloat(tagValue, 64)))
		case reflect.Int:
			v.SetInt(int64(xerr.Must(strconv.Atoi(tagValue))))
		case reflect.Slice:
			strs := strings.Split(tagValue, ",")
			parseStrsToSlice(t, v, strs)
		}
	}
}

func New[T any]() *T {
	v := new(T)
	xref.WalkValues(v, walkToSetDefaultValues)
	return v
}
