package reflection

import (
	"reflect"
)

// 编写函数 walk(x interface{}, fn func(string))，参数为结构体 x，并对 x 中的所有字符串字段调用 fn 函数
func walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Interface:
			walk(field.Interface(), fn)
		}
	}
}
