package main

import "fmt"

//121. 买卖股票的最佳时机
//url: leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
func main() {
	var stock = []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(stock)
	fmt.Println(result)
}

func maxProfit(prices []int) int {
	//初始化最大收益
	maxProfit := 0
	//最低买入值
	minBuy := prices[0]

	if len(prices) <= 1 {
		return maxProfit
	}

	for i := 1; i < len(prices); i++ {
		if minBuy > prices[i] {
			minBuy = prices[i]
		} else {
			sellVal := prices[i] - minBuy
			maxProfit = max(sellVal, maxProfit)
		}
	}
	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
