package _14_Best_Time_to_Buy_and_Sell_Stock_with_Transaction_Fee

func maxProfit(prices []int, fee int) int {
	if prices == nil || len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	MP := make([][]int, 2, 2)
	MP[0] = make([]int, len(prices), len(prices))
	MP[1] = make([]int, len(prices), len(prices))

	MP[0][0] = 0
	MP[1][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		MP[0][i] = Max(MP[0][i-1], MP[1][i-1]+prices[i]-fee)
		MP[1][i] = Max(MP[0][i-1]-prices[i], MP[1][i-1])
	}
	return MP[0][len(prices) - 1]
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