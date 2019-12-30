package main

import (
	"fmt"
	"go/build"
	"path/filepath"
)

var x = build.NoGoError{"wawa"}
var y = filepath.SkipDir

type eror interface {
	Eror() string
}

func New(text string) eror {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Eror() string {
	return e.s
}


func main() {
	x := errorString{"text"}
	a := New("牛逼23")
	fmt.Println(a)
	fmt.Println(x)
	fmt.Println(a.Eror())
}
