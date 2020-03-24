package main

import "fmt"

func mapToAnotherFunction(m map[string]int) {
	m["hello"] = 3
	m["world"] = 4
	m["new_word"] = 5
	m = make(map[string]int)  //这说明m是指针, 而非c++含义上的引用
}

// func mapToAnotherFunctionAsRef(m *map[string]int) {
// m["hello"] = 30
// m["world"] = 40
// m["2ndFunction"] = 5
// }

func main() {
	m := make(map[string]int)
	m["hello"] = 1
	m["world"] = 2

	// Initial State
	for key, val := range m {
		fmt.Println(key, "=>", val)
	}

	fmt.Println("-----------------------")

	mapToAnotherFunction(m)
	// After Passing to the function as a pointer
	for key, val := range m {
		fmt.Println(key, "=>", val)
	}
	v, ok := m["a"]
	fmt.Println(v, ok)
	// Try Un Commenting This Line
	fmt.Println("-----------------------")

	// mapToAnotherFunctionAsRef(&m)
	// // After Passing to the function as a pointer
	// for key, val := range m {
	//  fmt.Println(key, "=>", val)
	// }

	// Outputs
	// prog.go:12:4: invalid operation: m["hello"] (type *map[string]int does not support indexing)
	// pr
	//og.go:13:4: invalid operation: m["world"] (type *map[string]int does not support indexing)
	// prog.go:14:4: invalid operation: m["2ndFunction"] (type *map[string]int does not support indexing)

}