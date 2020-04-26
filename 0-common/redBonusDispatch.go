package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	result := RedBonusDispatch(100, 3)
	fmt.Println(result)
}

//红包分配算法
//totalMoney float64 红包总额
//n int 红包个数
func RedBonusDispatch(totalMoney float64, n int) []float64 {
	rand.Seed(time.Now().UnixNano())

	//单个红包基准金额
	var base = 30.01

	//基准总额
	baseTotal := base * float64(n)

	//剩余总额
	remainTotal := totalMoney - float64(baseTotal)

	remainItems := make([]float64, 0)

	var randSum float64 = 0
	for i := 0; i < n; i++ {
		randNum := float64(rand.Intn(100))
		remainItems = append(remainItems, float64(randNum))
		randSum += randNum
	}

	remainItemMoney := make([]float64, 0)
	var remainItemMoneySum float64 = 0
	for _, val := range remainItems {
		//通过随机数占比获取红包金额
		moneyEvery := (val / randSum) * remainTotal
		moneyEvery += base

		//四舍五入
		money, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", moneyEvery), 64)
		remainItemMoneySum += money
		remainItemMoney = append(remainItemMoney, money)
	}

	remainItemMoneySum, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", remainItemMoneySum), 64)

	diffMoney := totalMoney - remainItemMoneySum

	remainItemMoney[0] += diffMoney

	firstMoneyDetail, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", remainItemMoney[0]), 64)
	remainItemMoney[0] = firstMoneyDetail

	var sum float64 = 0
	for _, i2 := range remainItemMoney {
		sum += i2
	}
	return remainItemMoney

}
