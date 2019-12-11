package main

import "fmt"

type ContainCanUseFaceID interface {
	CanUseFaceID()
}

type ContainStolen interface {
	Stolen()
}

type Alipay struct {
}

func (a *Alipay) CanUseFaceID() {

}

type Cash struct {
}

func (c *Cash) Stolen() {

}

func Print(payMethod interface{}) {
	fmt.Printf("%T\n", payMethod)
	switch payMethod.(type) {
	case ContainCanUseFaceID:
		fmt.Println("ContainCanUseFaceID")
	case ContainStolen:
		fmt.Println("ContainStolen")
	}
}

func main() {
	Print(new(Alipay))
	Print(new(Cash))
}
