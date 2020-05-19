package main

import "fmt"

//152. 乘积最大子数组
//leetcode-cn.com/problems/maximum-product-subarray/
// 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
func main() {
	var nums = make([]int, 0)
	nums = []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	product := nums[0]
	minPro := nums[0]
	maxPro := nums[0]

	for i := 1; i < len(nums); i++ {
		maxVal := maxPro
		minVal := minPro

		////分别假设nums[i] 是正数，负数
		maxPro = GetMax(maxVal*nums[i], GetMax(minVal*nums[i], nums[i]))
		minPro = GetMin(minVal*nums[i], GetMin(maxVal*nums[i], nums[i]))

		product = GetMax(maxPro, product)
	}
	return product
}

//获取两个数字的最大值
func GetMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//获取两个数的最小值
func GetMin(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
