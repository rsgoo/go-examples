package main

import (
	"os"
	"fmt"
)

func main() {
	//tmp := [1]byte{}
	s := make([]byte, 1024)
	//n, err := os.Stdin.Read(tmp[:])
	n, err := os.Stdin.Read(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(s))

	fmt.Println(n)

	v := struct{}{}

	fmt.Printf("%T", v)
}
