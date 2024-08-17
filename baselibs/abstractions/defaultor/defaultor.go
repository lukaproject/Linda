package defaultor

import (
	"Linda/baselibs/abstractions/xref"
	"reflect"
	"strconv"
	"strings"

	"github.com/ecodeclub/ekit/slice"
	"github.com/lukaproject/xerr"
)

var (
	defaultorTagKey = "xdefault"
)

func parseStrsToSlice(t reflect.Type, v reflect.Value, strs []string) {
	switch t.Elem().Kind() {
	case reflect.String:
		v.Set(reflect.ValueOf(strs))
	case reflect.Int:
		v.Set(
			reflect.ValueOf(
				slice.Map(strs, func(idx int, str string) int {
					return int(xerr.Must(strconv.ParseInt(str, 10, 64)))
				})))
	case reflect.Float32:
		v.Set(
			reflect.ValueOf(
				slice.Map(strs, func(idx int, str string) float32 {
					return float32(xerr.Must(strconv.ParseFloat(str, 32)))
				})))
	case reflect.Float64:
		v.Set(
			reflect.ValueOf(
				slice.Map(strs, func(idx int, str string) float64 {
					return xerr.Must(strconv.ParseFloat(str, 64))
				})))
	default:
	}
}

func walkToSetDefaultValues(input xref.WalkFuncInput) {
	tags, t, v := input.FieldTag, input.Type, input.Value
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
		case reflect.Bool:
			v.SetBool(xerr.Must(strconv.ParseBool(tagValue)))
		}
	}
}

func New[T any]() *T {
	v := new(T)
	xref.WalkValues(v, walkToSetDefaultValues)
	return v
}
