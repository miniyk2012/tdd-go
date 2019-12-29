package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

type Circle struct {
	Radius float64
}

// 在 Go 语言中 interface resolution 是隐式的。如果传入的类型匹配接口需要的，则编译正确
func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}





