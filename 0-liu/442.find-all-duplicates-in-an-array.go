package main

import (
	"sort"
)

//442. 数组中重复的数据
/*给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。
 *找到所有出现两次的元素。
 *你可以不用到任何额外空间并在O(n)时间复杂度内解决这个问题吗？
 */
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-all-duplicates-in-an-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	var nums = []int{4, 3, 2, 7, 8, 2, 3, 1}
	findDuplicates(nums)
	findDuplicatesBySort(nums)
}

//通过排序的方式
func findDuplicatesBySort(nums []int) []int {
	var result = make([]int, 0)
	count := len(nums)
	if count < 1 {
		return result
	}
	sort.Ints(nums)
	//fmt.Println(nums)
	for i := 1; i < count; i++ {
		if nums[i]-nums[i-1] == 0 {
			result = append(result, nums[i])
		}
	}
	//fmt.Println(result)
	return result
}

//使用maps标识的方式
func findDuplicates(nums []int) []int {
	var result = make([]int, 0)
	count := len(nums)
	if count < 1 {
		return result
	}
	maps := make(map[int]int, count)

	for _, val := range nums {
		_, ok := maps[val]
		if !ok {
			maps[val] = 1
		} else {
			maps[val]++
			result = append(result, val)
		}
	}
	//fmt.Println(maps)
	//fmt.Println(result)

	return result
}
