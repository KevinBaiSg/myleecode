package _123_Best_Time_to_Buy_and_Sell_Stock_III

import "math"

func maxProfit(k int, prices []int) int {
	if k == 0 { return 0 }
	if prices == nil || len(prices) == 0 || len(prices) == 1 { return 0 }

	res := 0
	// 特殊处理
	if k > len(prices) / 2 {
		var (
			MP0 = 0
			MP1 = -prices[0]
		)
		for i := 1; i < len(prices); i++ {
			MP0, MP1 = Max(MP0, MP1 + prices[i]), Max(MP0 - prices[i], MP1)
			res = Max(res, MP0, MP1)
		}
		return res
	}

	// MP[i][x][j] j=0:没有股票 x:交易次数 j=1:拥有一股
	MP := make([][][]int, len(prices), len(prices))
	for i := 0; i < len(prices); i++ {
		MP[i] = make([][]int, k+1, k+1)
		for x := 0; x <= k; x++ {
			MP[i][x] = make([]int, 2, 2)
		}
	}

	for x := 0; x <= k ; x++ {
		if x == 0 {
			MP[0][x][0] = 0
			MP[0][x][1] = -prices[0]
		} else {
			MP[0][x][0], MP[0][x][1] = math.MinInt32, math.MinInt32
		}
	}

	for i := 1; i < len(prices); i++  {
		for x := 0; x <= k; x++  {
			if x == 0 {
				MP[i][0][0] = MP[i-1][0][0]
				MP[i][0][1] = Max(MP[i-1][0][0] - prices[i], MP[i-1][0][1])
				res = Max(res, MP[i][0][0], MP[i][0][1])
			} else if x == k {
				MP[i][x][0] = Max(MP[i-1][x][0], MP[i-1][x-1][1]+prices[i])
				res = Max(res, MP[i][x][0])
			} else {
				MP[i][x][0] = Max(MP[i-1][x][0], MP[i-1][x-1][1]+prices[i])
				MP[i][x][1] = Max(MP[i-1][x][0]-prices[i], MP[i-1][x][1])
				res = Max(res, MP[i][x][0], MP[i][x][1])
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

//func MaxRes(mp [][]int, k int) int {
//	max := mp[0][0]
//	for i := 1; i < k ; i++ {
//		if mp[i][0] > max {
//			max = mp[i][0]
//		}
//	}
//	return max
//}