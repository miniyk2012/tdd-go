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
	copy(s, t)  // 拷贝len(s)长度的数据
	return t
}


