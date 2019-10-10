package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
	OS string
}

func (nokiaPhone *NokiaPhone) call() {
	fmt.Println("nokia os is ", nokiaPhone.OS)
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
	OS string
}

func (iPhone *IPhone) call() {
	fmt.Println("iPhone os is ", iPhone.OS)
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone
	phone = &NokiaPhone{OS:"WP"}
	phone.call()

	phone = new(IPhone)
	phone = &IPhone{OS:"iOS"}
	phone.call()

}
