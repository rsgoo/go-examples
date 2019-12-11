package main

import (
	"sync"
	"fmt"
)

var (
	count      int
	countGuard sync.Mutex
)

func main() {
	SetCount(1)
	fmt.Println(GetCount())
}

func GetCount() int {
	countGuard.Lock()
	defer countGuard.Unlock()
	return count
}

func SetCount(n int) {
	countGuard.Lock()
	count = n
	countGuard.Unlock()
}
