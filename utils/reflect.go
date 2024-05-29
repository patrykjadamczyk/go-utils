package utils

import (
	"reflect"
)

// Check if specified value is array
func IsArray(val any) bool {
	kind := reflect.TypeOf(val).Kind()
	return kind == reflect.Array || kind == reflect.Slice
}

// Check if specified value is function
func IsFunc(val any) bool {
	kind := reflect.TypeOf(val).Kind()
	return kind == reflect.Func
}

// Call specified function with specified arguments
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

// Call specified function without arguments
func CallFuncWithoutArgs(val any) []reflect.Value {
	function := reflect.ValueOf(val)
	rargs := []reflect.Value{}
	if function.Kind() != reflect.Func {
		panic("Value is not a function")
	}
	return function.Call(rargs)
}

// Change specified reflection values to values of type any
func ReflectValueToValue(values ...reflect.Value) []any {
	result := []any{}
	for _, v := range values {
		result = append(result, v.Interface())
	}
	return result
}

// Get Zero Value of specified type through reflection
func GetZeroReflectionValue[T any]() reflect.Value {
	var temp T
	return reflect.Zero(reflect.TypeOf(temp))
}

// Get Type Name using any
func GetTypeNameAny(v any) string {
	return GetTypeName(v)
}

// Get Type Name
func GetTypeName[T any](v T) string {
	return reflect.TypeOf(v).String()
}
