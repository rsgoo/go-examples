package main

import "fmt"

//448. 找到所有数组中消失的数字
//leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/
/** 给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
 *  找到所有在 [1, n] 范围之间没有出现在数组中的数字。
 */
func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}

func findDisappearedNumbers(nums []int) []int {
	var result = make([]int, 0)
	count := len(nums)
	maps := make(map[int]int, count)

	if count == 0 {
		return result
	}

	for _, val := range nums {
		maps[val] = 1
	}

	for i := 1; i <= count; i++ {
		if maps[i] == 0 {
			result = append(result, i)
		}
	}

	return result
}
