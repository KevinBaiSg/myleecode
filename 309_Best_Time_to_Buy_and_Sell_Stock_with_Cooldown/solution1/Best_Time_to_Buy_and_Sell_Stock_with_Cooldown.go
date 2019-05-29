package solution1

func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	MP := make([][][]int, len(prices), len(prices))
	for i := 0; i < len(prices); i++ {
		MP[i] = make([][]int, 2, 2)
		for j := 0; j < 2 ; j++ {
			MP[i][j] = make([]int, 2, 2)
		}
	}

	MP[0][0][0] = 0
	MP[0][1][0] = -prices[0]
	MP[0][0][1] = -1 << 32
	MP[0][1][1] = -1 << 32

	for i := 1; i < len(prices); i++ {
		MP[i][0][0] = Max(MP[i-1][0][0], MP[i-1][0][1])
		MP[i][0][1] = MP[i-1][1][0] + prices[i]
		MP[i][1][0] = Max(MP[i-1][0][0]-prices[i], MP[i-1][1][0])
		MP[i][1][1] = -1 << 32
	}
	end := len(prices) -1
	return Max(MP[end][0][0], MP[end][0][1], MP[end][1][0])
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