package main

import "fmt"

//167. 两数之和 II - 输入有序数组
/**
 * 给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
 * 函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
 *来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func main() {
	var arr = []int{2, 3, 6, 7, 11, 15}
	fmt.Println(twoSum(arr, 9))
}

func twoSum(numbers []int, target int) [][]int {
	result := [][]int{}
	left := 0
	right := len(numbers) - 1

	for left <= right {
		if numbers[left]+numbers[right] == target {
			temp := []int{left, right}
			result = append(result, temp)
			left++
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}
	return result
}
