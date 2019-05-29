package solution2

func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	MP := make([][][]int, len(prices), len(prices))


	MP[i,0,0]=Max(MP[i-1,0,0], MP[i-1,0,1])
	MP[i,0,1]=MP[i,1,0]+prices[i]
	MP[i,1,0]=Max(MP[i,0,0]-prices[i],MP[i,1,0])
	res = Max(MP[i,0,0], MP[i,0,1], MP[i,1,0])
}