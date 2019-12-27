package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%v,%T\n", time.April, time.April)
	fmt.Println(time.Now().Month().String() == "December")

	//自定义一个时间戳
	timeObj := time.Unix(1577749841, 0)

	fmt.Println(timeObj.Month())
	fmt.Println(timeObj.YearDay())
	fmt.Println(timeObj.Format("2006-01-02 15:04:05:999"))

	baseNum := 10000000000

	//1s=1000ms，1 ms=1000μs，1μs=1000ns
	sum := 0
	for i := 1; i <= baseNum; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)
	fmt.Println("开始：", now)
	fmt.Println("结束：", time.Now())
	fmt.Println("耗时：", time.Since(now))

	fmt.Println("-------------------------------")

	newNow := time.Now()
	newSum := (1 + baseNum) * baseNum / 2

	fmt.Println("sum = ", newSum)
	fmt.Println("new 开始：", newNow)
	fmt.Println("new 结束：", time.Now())
	fmt.Println("new 耗时：", time.Since(newNow))

}
