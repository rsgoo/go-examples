package main

import (
	"fmt"
	"sync"
)

var count int
var countMux sync.RWMutex

func GetCount() int {
	countMux.RLock()
	defer countMux.RUnlock()
	return count
}

func IncCount() {
	countMux.Lock()
	defer countMux.Unlock()
	count++
}

func main() {
	IncCount()
	count := GetCount()
	fmt.Println(count)
}
