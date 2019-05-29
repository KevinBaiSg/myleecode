
### 问题描述
```
123. Best Time to Buy and Sell Stock III
Hard

Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most two transactions.

Note: You may not engage in multiple transactions at the same time 
(i.e., you must sell the stock before you buy again).

Example 1:

Input: [3,3,5,0,0,3,1,4]
Output: 6
Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
             Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
Example 2:

Input: [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
             Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
             engaging multiple transactions at the same time. You must sell before buying again.
Example 3:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
```  

### 解决方法

#### Solution1
* 思路 DP
MaxProfit=MP 最大收益

1.状态定义：MP[i,j,k]    到i天最大收益
                        j=0:没有股票 j=1:拥有一股 k:交易次数
                             
2.  for i
        for k
            MP[i,k,0]=Max(MP[i-1,k,0],MP[i-1,k,1]+prices[i])
            MP[i,k,1]=Max(MP[i-1,k-1,0]-prices[i],MP[i-1,k,1])
            
#### Solution2
* 思路 DP
通过数组记录每次buy & sell 后的最大值（eg： 第一次买入3元， 其实就是目前收入 -3）。

1.第k次buy就是 上一次卖出剩下的钱 - 本次购买需要的钱
2.第k次sell就是 第k次买入 + 本次卖出的钱

 buy[j] = max(buy[j], sell[j-1] - prices[i])
 sell[j] = max(sell[j], buy[j] + prices[i])
 
 
 ### 优化
 如果k > len(prices)/2 不用考虑交易次数，可以把问题转化成无限购买次数