package main

import (
	"fmt"
)

// this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func main() {
	n := 0
	reply := &n
	Multiply(10, 5, reply)

	//Multiply: 0xc00004a058
	fmt.Println("Multiply:", reply)

	//Multiply: 0xc000072018
	fmt.Println("Multiply:", &reply)

	//ultiply: 50
	fmt.Println("Multiply:",*reply)

	//Multiply: 0xc00004a058
	fmt.Println("Multiply:",*(&reply))

	//Multiply: 0xc00004a058
	fmt.Println("Multiply:",&(*reply))

	//Multiply: 50
	fmt.Println("Multiply:",*&(*reply)) // Multiply: 50
}
