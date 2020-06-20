package main

import (
	"fmt"
	"sort"
)

//4. 寻找两个正序数组的中位数
/**
 * 给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
 * 请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
 * 你可以假设 nums1 和 nums2 不会同时为空。
 */
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	fmt.Println(findMedianSortedArraysV2(nums1, nums2))

}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	fmt.Println(nums1)

	//Go中 3/2 = 1； 4/2=2
	commonMid := len(nums1) / 2
	if len(nums1)%2 == 0 {
		//fmt.Println(commonMid)
		//fmt.Println(commonMid - 1)
		return (float64(nums1[commonMid]) + float64(nums1[commonMid-1])) / 2
	} else {
		return float64(nums1[commonMid])
	}
}
