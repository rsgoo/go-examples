package main

import (
	"fmt"
)

type Phone interface {
	call()
	sell()
}

type NokiaPhone struct {
	OS string
}

func (nokiaPhone *NokiaPhone) call() {
	fmt.Println("nokia os is ", nokiaPhone.OS)
	fmt.Println("I am Nokia, I can call you!")
}

func (nokiaPhone *NokiaPhone) sell() {
	fmt.Println("sell nokia")
}

type IPhone struct {
	OS string
}

func (iPhone *IPhone) call() {
	fmt.Println("iPhone os is ", iPhone.OS)
	fmt.Println("I am iPhone, I can call you!")
}

func (iPhone *IPhone) sell() {
	fmt.Println("sell iPhone")
}

func main() {
	var nokia Phone
	var iPhone Phone

	nokia = &NokiaPhone{OS: "WP"}
	nokia.call()
	nokia.sell()

	iPhone = &IPhone{OS: "iOS"}
	iPhone.call()
	iPhone.sell()

}
