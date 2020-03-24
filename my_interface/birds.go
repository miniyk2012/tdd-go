package main

import (
	"fmt"
	"reflect"
)

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
	fmt.Printf("ChickenAnimation of %T\n", chicken)
	chicken.Walk()
	println(chicken.Twitter())
}

func NilInterfaceTest(chicken Chicken) {
	if chicken == nil || reflect.ValueOf(chicken).IsNil() {
		fmt.Println("Sorry,It’s Nil")
	} else {
		fmt.Println("Animation Start!")
		fmt.Printf("type:%v,value:%v\n", reflect.TypeOf(chicken), reflect.ValueOf(chicken))
		ChickenAnimation(chicken)
	}
}
type TestInt int

func main() {
	var bird Birds
	sparrow := &Sparrow{name: "yangkai"}
	bird = sparrow
	BirdAnimation(bird, 1000)
	// 或者将sparrow直接作为参数
	BirdAnimation(sparrow, 1000)

	println()
	var sparrow3 *Sparrow
	NilInterfaceTest(sparrow3)

	println()
	var chicken Chicken
	sparrow2 := Sparrow{name: "ouyang"}
	chicken = &sparrow2
	/*
	Chicken接口的Walk方法的接收者是非指针的Sparrow，我们把&Sparrow赋值给Chicken接口变量为什么可以通过?

	首先，一个指针类型的方法列表必然包含所有接收者为指针接收者的方法，同理非指针类型的方法列表也包含所有接收者为非指针类型的方法。
	在我们例子中*Sparrow首先包含：Fly和Twitter；Sparrow包含Walk。

	其次，当我们拥有一个指针类型的时候，因为有了这个变量的地址，我们得到这个具体的变量，所以一个指针类型的方法列表还可以包含其非指针类型作为接收者的方法。
	在我们的例子中就是*Sparrow的方法列表为：Fly、Twitter和Walk，所以chicken = &sparrow2可以通过。
	 */
	ChickenAnimation(chicken)

}
