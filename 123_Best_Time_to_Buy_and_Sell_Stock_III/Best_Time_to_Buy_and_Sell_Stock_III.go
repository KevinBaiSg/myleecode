package _123_Best_Time_to_Buy_and_Sell_Stock_III

import "math"

func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 || len(prices) == 1 { return 0 }

	res := 0
	const K = 2 + 1
	// MP[i][k][j] j=0:没有股票 k:交易次数 j=1:拥有一股
	MP := make([][][]int, len(prices), len(prices))
	for i := 0; i < len(prices); i++ {
		MP[i] = make([][]int, K, K)
		for k := 0; k < K; k++ {
			MP[i][k] = make([]int, 2, 2)
		}
	}

	for k := 0; k < K ; k++ {
		if k == 0 {
			MP[0][k][0] = 0
			MP[0][k][1] = -prices[0]
		} else {
			MP[0][k][0], MP[0][k][1] = math.MinInt32, math.MinInt32
		}
	}

	for i := 1; i < len(prices); i++  {
		for k := 0; k < K; k++  {
			if k == 0 {
				MP[i][0][0] = MP[i-1][0][0]
				MP[i][0][1] = Max(MP[i-1][0][0] - prices[i], MP[i-1][0][1])
				res = Max(res, MP[i][0][0], MP[i][0][1])
			} else if k == K {
				MP[i][k][0] = Max(MP[i-1][k][0], MP[i-1][k-1][1]+prices[i])
				res = Max(res, MP[i][k][0])
			} else {
				MP[i][k][0] = Max(MP[i-1][k][0], MP[i-1][k-1][1]+prices[i])
				MP[i][k][1] = Max(MP[i-1][k][0]-prices[i], MP[i-1][k][1])
				res = Max(res, MP[i][k][0], MP[i][k][1])
			}
		}
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