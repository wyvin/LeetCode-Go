package _0121_Best_Time_to_Buy_and_Sell_Stock

import "math"

// 2021-01-31
// 一次遍历：
// 每次循环得到左边最大差值和最低点，注意求最值得顺序不能调转
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	max := 0
	min := math.MaxInt32
	for _, p := range prices {
		max = Max(p-min, max)
		min = Min(min, p)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
