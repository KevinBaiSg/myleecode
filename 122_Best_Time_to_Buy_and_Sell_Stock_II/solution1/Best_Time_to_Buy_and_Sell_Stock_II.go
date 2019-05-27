package solution1

func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 || len(prices) == 1 { return 0 }

	res := 0
	for i := 1; i < len(prices); i++  {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}	
	}
	return res
}