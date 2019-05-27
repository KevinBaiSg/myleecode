package solution3

func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 || len(prices) == 1 { return 0 }

	var (
		res = 0
		MP0 = 0
		MP1 = -prices[0]
	)
	for i := 1; i < len(prices); i++ {
		MP0, MP1 = Max(MP0, MP1 + prices[i]), Max(MP0 - prices[i], MP1)
		res = Max(res, MP0, MP1)
	}
	return res
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