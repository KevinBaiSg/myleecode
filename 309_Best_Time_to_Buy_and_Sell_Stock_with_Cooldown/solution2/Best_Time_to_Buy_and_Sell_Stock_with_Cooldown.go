package solution2

func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	buy := make([]int, len(prices), len(prices))
	sell := make([]int, len(prices), len(prices))
	cooldown := make([]int, len(prices), len(prices))

	buy[0] = -prices[0]
	sell[0] = 0
	cooldown[0] = 0

	for i := 1; i < len(prices); i++ {
		buy[i] = Max(buy[i-1], cooldown[i-1]-prices[i])
		sell[i] = Max(buy[i-1]+prices[i], sell[i-1])
		cooldown[i] = sell[i-1]
	}
	return Max(sell[len(prices) - 1], cooldown[len(prices) - 1])
}

func Max(m ...int) int {
	if len(m) == 0 { return 0 }
	max := m[0]
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}