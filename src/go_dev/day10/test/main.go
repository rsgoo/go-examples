package main

import (
	"reflect"
	"fmt"
)

func main() {
	var str = "hello"
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("world")
	fmt.Printf("%v\n", fs)		//0xc42000e1e0
	fmt.Printf("%v\n", str)		//world
}
