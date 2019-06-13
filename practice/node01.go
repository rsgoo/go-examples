package main

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func main() {
	fmt.Println("plz enter your firstName")
	fmt.Scanln(&firstName)
	fmt.Println("plz enter your lastName")
	fmt.Scanln(&lastName)
	fmt.Printf("Hi, %s %s\n", firstName, lastName)
	fmt.Println()
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println(f)
	fmt.Println(i)
	fmt.Println(s)
}
