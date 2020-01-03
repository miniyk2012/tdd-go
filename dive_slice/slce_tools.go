package dive_slice

import (
	"fmt"
	"regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")
func AppendByte(slice []byte, data ...byte) []byte {
	return nil
}



func Append(s []int, v int) {
	a := append(s, v)
	fmt.Printf("a.cap=%d, a.len=%d, a=%v\n", cap(a), len(a), a)
	fmt.Printf("s.cap=%d, s.len=%d, s=%v\n", cap(s), len(s), s)
	fmt.Printf("%p\n", &s)
}
