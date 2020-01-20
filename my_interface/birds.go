package main

import "fmt"

type Birds interface {
	Twitter() string
	Fly(high int) bool
}

type Chicken interface {
	Birds // 继承了Birds的2个方法
	Walk()
}

type Sparrow struct {
	name string
}

// 指针方法
func (s *Sparrow) Fly(hign int) bool {
	// ...
	return true
}

// 指针方法
func (s *Sparrow) Twitter() string {
	// ...
	return fmt.Sprintf("%s,jojojo", s.name)
}

// 值方法
func (s Sparrow) Walk() {
	println("Walk")
}

func BirdAnimation(bird Birds, high int) {
	fmt.Printf("BirdAnimation of %T\n", bird)
	println(bird.Twitter())
	bird.Fly(high)
}

func ChickenAnimation(chicken Chicken) {
	fmt.Printf("BirdAnimation of %T\n", chicken)
	chicken.Walk()
}

func NilInterfaceTest(chicken Chicken) {
	if chicken == nil {
		fmt.Println("Sorry,It’s Nil")
	} else {
		fmt.Println("Animation Start!")
		ChickenAnimation(chicken)
	}
}


func main() {
	var bird Birds
	sparrow := &Sparrow{name: "yangkai"}
	bird = sparrow
	BirdAnimation(bird, 1000)
	// 或者将sparrow直接作为参数
	BirdAnimation(sparrow, 1000)

	var sparrow3 *Sparrow
	NilInterfaceTest(sparrow3)
}
