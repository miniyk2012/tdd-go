package dive_slice

import (
	"regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

func AppendByte(slice []byte, data ...byte) []byte {
	return nil
}

// 扩Cap, 则需要新建一个slice, 把旧的复制过去
func DoubleCap(s []string) []string {
	t := make([]string, len(s), (cap(s)*2)+1)
	//for i, v := range s {
	//	t[i] = v
	//}
	copy(t, s)  // 拷贝len(s)长度的数据
	return t
}

func AppendString(slice []string, data ...string) []string {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]string, (n+1) * 2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:n]  // 缩短至长度n
	copy(slice[m:n], data)
	return slice
}


func Filter(s []int, fn func(int) bool) []int {
	var r []int  // nil
	for i := range s {
		if fn(s[i]) {
			r = append(r, s[i])
		}
	}
	return r
}

