package main

import "fmt"

func main() {
	var successSlice = []int{88243, 84470, 82350, 64733, 5368}
	successCount := GetSum(successSlice)
	fmt.Println(successCount)
	var allStatistics = []int{successCount, 73467, 11794}
	var allCount = GetSum(allStatistics)
	fmt.Println(allCount)
	fmt.Println(float64(successCount) / float64(allCount) * 750)

	fmt.Println(266 * 13)
	fmt.Println(60 * 24)

	//var cacheSlice = []int{1311012, 648602, 31215, 658786, 13424}
	var cacheSlice = []int{980000, 648602, 31215, 658786, 13424}
	var needCacheCount = 2663041
	var cacheSliceCount = GetSum(cacheSlice);
	fmt.Println("needCacheCount :", needCacheCount)
	fmt.Println("cacheSliceCount:", cacheSliceCount)
	fmt.Println("cacheSliceDiff :", needCacheCount-cacheSliceCount)
	fmt.Println(needCacheCount * 3.0 /1024.0/1024.0)

}

func GetSum(slices []int) (sum int) {
	for _, value := range slices {
		sum += value
	}
	return sum
}
