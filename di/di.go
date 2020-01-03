package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// 用 writer 来把name发送到某处, 依赖注入
func Greet(writer io.Writer, name string) {
	// Fprintf formats according to a format specifier and writes to w.
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	Greet(os.Stdout, "Elodie")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
	Greet(os.Stdout, "Yangkai")  // 不会运行到
}