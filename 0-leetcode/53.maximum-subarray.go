package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	rest := math.MinInt32
	maxSum := 0
	for i := 0; i < len(nums); i++ {
		maxSum += nums[i]
		rest = getMax(maxSum, rest)

		//这里一定要置为0
		if maxSum < 0 {
			maxSum = 0
		}
	}
	return rest
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
