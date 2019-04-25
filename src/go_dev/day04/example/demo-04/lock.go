package demo_04

import (
	"math/rand"
	"fmt"
	"time"
	"sync"
)

var Lock sync.Mutex

func TestMapLock() {
	rand.Seed(time.Now().UnixNano())
	var a map[int]int
	a = make(map[int]int, 5)
	a[0] = 1
	a[2] = 22
	a[3] = 333
	a[4] = 4444
	a[5] = 55555
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			Lock.Lock()
			b[0] = rand.Intn(100)
			Lock.Unlock()
		}(a)
	}
	Lock.Lock()
	fmt.Println(a)
	Lock.Unlock()
}
