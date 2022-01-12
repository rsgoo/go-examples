package main

import "fmt"

type IAnimal interface {
	Jump()
}

type IAction interface {
	Jump()
}

type Dog struct {
}

func (this *Dog) Jump() {
	fmt.Println("dog jump")
}

func main() {
	var animal IAnimal = &Dog{}

	//animal 对象是否实现了 IAction 接口
	action, ok := animal.(IAction)
	if ok {
		action.Jump()
	}
}
