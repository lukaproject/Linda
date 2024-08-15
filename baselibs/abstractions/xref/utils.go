package xref

import (
	"reflect"
)

// walkFuncType
// params:
// for each one reflect items, the params's means:
// fieldName 	string
// fieldTag 	reflect.StructTag
// type 		reflect.Type
// value 		reflect.Value
type walkFuncType = func(string, reflect.StructTag, reflect.Type, reflect.Value)

// WalkVaules
// 递归遍历x及其参数的所有参数，对每一个参数执行walkFunc
func WalkValues(x any, walkFunc walkFuncType) {
	xtypes := reflect.TypeOf(x)
	xvalues := reflect.ValueOf(x)
	if xvalues.Kind() == reflect.Ptr && !xvalues.IsNil() {
		walkValuesImpl("", "", xtypes.Elem(), xvalues.Elem(), walkFunc)
	}
}

func walkValuesImpl(fieldName string, tags reflect.StructTag, t reflect.Type, v reflect.Value, walkFunc walkFuncType) {
	walkFunc(fieldName, tags, t, v)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			v.Set(reflect.New(t.Elem()))
		}
		walkValuesImpl(fieldName, tags, t.Elem(), v.Elem(), walkFunc)
	}
	if v.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < v.NumField(); i++ {
		walkValuesImpl(t.Field(i).Name, t.Field(i).Tag, t.Field(i).Type, v.Field(i), walkFunc)
	}
}
