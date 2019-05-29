package _123_Best_Time_to_Buy_and_Sell_Stock_III

func maxProfit(k int, prices []int) int {
	if k >= len(prices)/2 {
		ret := 0
		for i := 1; i < len(prices); i++ {
			if v := prices[i] - prices[i-1]; v > 0 {
				ret += v
			}
		}
		return ret
	}
	buy := make([]int, k + 1)
	sell := make([]int, k + 1)
	for i := 0; i <= k; i++ {
		buy[i] = -2 << 31
	}
	for i := 0; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			buy[j] = max(buy[j], sell[j-1] - prices[i])
			sell[j] = max(sell[j], buy[j] + prices[i])
		}
	}
	return sell[k]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}