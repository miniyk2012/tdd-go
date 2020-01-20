package reflection

import (
	"reflect"
)

// 编写函数 walk(x interface{}, fn func(string))，参数为结构体 x，并对 x 中的所有字符串字段调用 fn 函数
func walk(x interface{}, fn func(string)) {
	val := getValue(x)

	if val.Kind() == reflect.Slice {

	}
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i:=0; i<val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
