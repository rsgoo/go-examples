package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5}
	sum, pro := plusNumbs(nums...)
	fmt.Println(sum)
	fmt.Println(pro)
}

func plusNumbs(nums ...int) (sum, product int) {
	product = 1
	for _, value := range nums {
		sum += value
		product *= value
	}
	return sum, product
}
