package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {

	//fmt.Println("plz enter your firstName")
	//	fmt.Scanln(&firstName)
	//	fmt.Println("plz enter your lastName")
	//	fmt.Scanln(&lastName)
	//	fmt.Printf("Hi, %s %s\n", firstName, lastName)

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("plz enter some content:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("err is: ", err.Error())
	}
	fmt.Println("content is:", input)
}
