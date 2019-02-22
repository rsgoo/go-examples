package demo

import (
	"math"
	"fmt"
)

//水仙花
func GetFlowers() {

	for i := 0; i < 1000; i++ {

		var single = (i / 1) % 10 //个位
		var ten = (i / 10) % 10
		var hundred = (i / 100) % 10
		H3 := int(math.Pow(float64(hundred), float64(3)))
		T3 := int(math.Pow(float64(ten), float64(3)))
		S3 := int(math.Pow(float64(single), float64(3)))
		if H3+T3+S3 == i {
			fmt.Println("flower = ", i)
		}
	}

}
