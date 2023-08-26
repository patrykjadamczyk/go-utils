package utils

import (
	"reflect"
)

func IsArray(val any) bool {
	kind := reflect.TypeOf(val).Kind()
    return kind == reflect.Array || kind == reflect.Slice
}

func IsFunc(val any) bool {
    kind := reflect.TypeOf(val).Kind()
    return kind == reflect.Func
}

func CallFunc(val any, args ...any) []reflect.Value {
	function := reflect.ValueOf(val)
	rargs := []reflect.Value{}
	var errorType error
	for _, v := range args {
		switch v.(type) {
		case error:
			if v != nil {
				rargs = append(rargs, reflect.ValueOf(v))
			} else {
				rargs = append(rargs, reflect.Zero(reflect.TypeOf(errorType)))
			}
		case reflect.Value:
			rargs = append(rargs, v.(reflect.Value))
		default:
			if v != nil {
				rargs = append(rargs, reflect.ValueOf(v))
			} else {
				rargs = append(rargs, reflect.Zero(reflect.TypeOf(v)))
			}
		}
	}
	if function.Kind() != reflect.Func {
		panic("Value is not a function")
	}
	return function.Call(rargs)
}

func CallFuncWithoutArgs(val any) []reflect.Value {
	function := reflect.ValueOf(val)
	rargs := []reflect.Value{}
	if function.Kind() != reflect.Func {
		panic("Value is not a function")
	}
	return function.Call(rargs)
}

func ReflectValueToValue(values ...reflect.Value) []any {
	result := []any{}
	for _, v := range values {
		result = append(result, v.Interface())
	}
	return result
}

func GetZeroReflectionValue[T any]() reflect.Value {
	var temp T
	return reflect.Zero(reflect.TypeOf(temp))
}
