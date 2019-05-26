package _21_Best_Time_to_Buy_and_Sell_Stock
/*
	121. Best Time to Buy and Sell Stock
	Easy

	Say you have an array for which the ith element is the price of a given stock on day i.

	If you were only permitted to complete at most one transaction
	(i.e., buy one and sell one share of the stock),
	design an algorithm to find the maximum profit.

	Note that you cannot sell a stock before you buy one.

	Example 1:
	Input: [7,1,5,3,6,4]
	Output: 5
	Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
				 Not 7-1 = 6, as selling price needs to be larger than buying price.
	Example 2:
	Input: [7,6,4,3,1]
	Output: 0
	Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/

/*
	解题思路：暴力求差
	计算在 i 卖出时，与其前边购买差最大值
*/

/*
	解题思路：DP
	状态定义：DP[i] 为当天卖出是最大利润
	DP状态转移：DP[i]=Max(maxProfit[i-1], prices[i]-minPrices)
*/
func maxProfit(prices []int) int {
	if prices == nil || len(prices) < 2 { return 0 }

	var (
		maxProfit = 0
		minPrices = prices[0]
	)
	
	for i := 1; i < len(prices); i++ {
		if prices[i] - minPrices < 0 {
			minPrices = prices[i]
		} else if prices[i] - minPrices > maxProfit {
			maxProfit = prices[i] - minPrices
		}
	}
	return maxProfit
}
