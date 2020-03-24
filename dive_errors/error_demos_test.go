package main

import (
	"container/list"
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	l1 := list.New()
	l1.PushBack(1)
	fmt.Println(l1.Back().Value)

	l2 := l1
	l2.PushBack(2)
	fmt.Println(l1.Back().Value, l2.Back().Value)
	fmt.Println()
	PrintList(l1)
	fmt.Println()
	PrintList(l2)

}

func TestType(t *testing.T) {
	var t1, t2 interface{} // t1, t2可以被赋值为任何对象
	t1, t2 = map[string]int{}, map[string]int{"a": 12}
	fmt.Print(t1, t2)
}

func PrintList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

type stringer interface {
	string() string
}

type tester interface {
	stringer // 嵌入其他接口
	test()
}

type data struct {
	z int
}

func (*data) test() {
	fmt.Print("test")
}

func (data) string() string {
	return "a"
}

func TestInterface(t *testing.T) {
	var d data  // 自动设置为zero值

	var a tester = &d
	fmt.Print(d)
	d.test()
	println(d.string())

	fmt.Print(a)
	a.test()
	println(a.string())
}


