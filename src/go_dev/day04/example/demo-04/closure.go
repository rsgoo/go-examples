package demo_04

import (
	"fmt"
	"strings"
)

func Adder() func(a int) int {
	var x int
	return func(a int) int {
		x += a
		return x
	}
}

func MakeSuffix(suffix string) func(name string) string {
	suffix = "." + suffix
	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}
}

func Closure() {
	var f = Adder()
	fmt.Println(f(1))   //1
	fmt.Println(f(10))  //11
	fmt.Println(f(100)) //111

	f1 := MakeSuffix("go")
	fmt.Println(f1("test"))
	f2 :=MakeSuffix("rs")
	fmt.Println(f2("rust.rs"))
	}
