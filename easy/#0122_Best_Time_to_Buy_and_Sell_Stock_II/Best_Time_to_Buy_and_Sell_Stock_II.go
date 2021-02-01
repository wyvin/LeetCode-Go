package _0122_Best_Time_to_Buy_and_Sell_Stock_II

// 2021-02-01
// 贪心，不在意操作过程，只在意总收益
func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	ans := 0
	for i := 1; i < len(prices); i++ {
		// price[i]>0
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}
