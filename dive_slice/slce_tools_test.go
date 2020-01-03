package dive_slice

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {
	s := make([]int, 0, 5)
	Append(s, 10)
	fmt.Printf("s.cap=%d, s.len=%d, s=%v\n", cap(s), len(s), s)
	fmt.Printf("%p", &s)
}