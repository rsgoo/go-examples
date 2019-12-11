package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getpid())
	fmt.Println(os.Getppid())
}
