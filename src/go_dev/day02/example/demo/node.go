package demo

import (
	"fmt"
	"math/rand"
	"time"
)

var A = "G"

func main() {
	rand.Seed(time.Now().UnixNano())
	n() //G
	m() //o
	n() //o
	fmt.Println(rand.Intn(10))
}

func n() {
	fmt.Println(A)
}

func m() {
	A = "o"
	fmt.Println(A)
}
