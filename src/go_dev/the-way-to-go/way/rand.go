package way

import (
	"math/rand"
	"fmt"
	"time"
)

func Rand() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
	fmt.Println()
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 5; i++ {
		r := rand.Intn(10)
		fmt.Printf("%d * ", r)
	}
	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%3.2f / ", 1000*rand.Float32())
	}
}