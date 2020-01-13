package dive_slice

import (
	"fmt"
	"reflect"
	"testing"
)

func AppendDemo(s []int, v int) {
	// s是个拷贝的slice, 但s底层的数组是同一个
	a := append(s, v)
	s = append(s, 2*v)
	s[2] = 100
	fmt.Printf("in function, a.cap=%d, a.len=%d, a=%v\n", cap(a), len(a), a)
	fmt.Printf("in function, s.cap=%d, s.len=%d, s=%v\n", cap(s), len(s), s)
	fmt.Printf("%p\n", &s)
}

func TestAppend(t *testing.T) {
	s := make([]int, 3, 5)
	AppendDemo(s, 10)
	fmt.Printf("s.cap=%d, s.len=%d, s=%v\n", cap(s), len(s), s)
	fmt.Printf("%p", &s)
}

func TestEqual(t *testing.T) {
	var s []byte
	fmt.Printf("%v\n", s)
	s = make([]byte, 4, 5)
	if !reflect.DeepEqual(s, []byte{0, 0, 0, 0}) {
		t.Error("xx")
	}

	x := [3]string{"Лайка", "Белка", "Стрелка"}
	x_s := x[:] // a slice referencing the storage of x
	x_s[1] = "12"
	fmt.Println(x_s)
	fmt.Println(x)
}

func TestCap(t *testing.T) {
	s := make([]byte, 5, 5)
	s = s[2:4]
	if len(s) != 2 || cap(s) != 3 {
		t.Error("xx")
	}
	s = s[:cap(s)]
	if len(s) != 3 || cap(s) != 3 {
		t.Error("xx")
	}
}

func TestGrowingSlice(t *testing.T) {
	s := []string{"a", "b"}
	new_s := DoubleCap(s)

	if len(new_s) != len(s) || cap(new_s) != 5 {
		t.Error("xx")
	}
	if !reflect.DeepEqual(s, new_s) || !reflect.DeepEqual(s, []string{"a", "b"}) {
		t.Error("xx")
	}

}

func TestAppendString(t *testing.T) {
	t.Run("需要扩容", func(t *testing.T) {
		s := make([]string, 2, 3)
		copy(s, []string{"a", "b"})
		s = AppendString(s, "c", "d", "e", "f", "g")
		if !reflect.DeepEqual(s, []string{"a", "b", "c", "d", "e", "f", "g"}) {
			t.Error("xx")
		}
		if cap(s) == 3 {
			t.Error("xx")
		}
	})

	t.Run("无需扩容", func(t *testing.T) {
		s := make([]string, 2, 100)
		copy(s, []string{"a", "b"})
		s = AppendString(s, "c", "d", "e", "f", "g")
		if !reflect.DeepEqual(s, []string{"a", "b", "c", "d", "e", "f", "g"}) {
			t.Error("xx")
		}
		if cap(s) != 100 {
			t.Error("xx")
		}
	})
}

func TestOriginAppend(t *testing.T) {
	a := make([]string, 2, 100)
	copy(a, []string{"John", "Paul"})
	b := []string{"George", "Ringo", "Pete"}
	a = append(a, b...)
	if !reflect.DeepEqual(a, []string{"John", "Paul", "George", "Ringo", "Pete"}) {
		t.Error("xx")
	}
	fmt.Println(cap(a))
}

func TestFilter(t *testing.T) {
	s := []int{5, 6, 7, 8, 8, 9, 100, -67, 0}
	fn := func(a int) bool {
		return a%2 == 0
	}
	p := Filter(s, fn)
	if !reflect.DeepEqual(p, []int{6,8,8,100,0}) {
		t.Error("xx")
	}
}

func TestFindDigits(t *testing.T) {
	filename := "file.txt"
	ret := FindDigits(filename)
	fmt.Printf("%s\n", ret)
	fmt.Println(string(ret))  // byte -> string
	if !reflect.DeepEqual(ret, []byte("943985")) {  // string -> byte
		t.Error("xxx")
	}
}
