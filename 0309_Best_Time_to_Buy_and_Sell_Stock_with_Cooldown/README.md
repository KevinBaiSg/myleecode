
### 问题描述
```
309. Best Time to Buy and Sell Stock with Cooldown
Medium

Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times) with the following restrictions:

You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)
Example:

Input: [1,2,3,0,2]
Output: 3 
Explanation: transactions = [buy, sell, cooldown, buy, sell]
```  

### 解决方法
#### solution1
* 思路 DP
MaxProfit=MP 最大收益

1.状态定义：MP[i,j,c]    到i天最大收益
                       j=0:没有股票 j=1:拥有一股 c: 是否为 cool down
                             
2.  for i
        MP[i,0,0]=Max(MP[i-1,0,0], MP[i-1,0,1])
        MP[i,0,1]=MP[i,1,0]+prices[i]
        MP[i,1,0]=Max(MP[i,0,0]-prices[i],MP[i,1,0])
        res = Max(MP[i,0,0], MP[i,0,1], MP[i,1,0])

#### solution2
* 思路 DP

#### Solution2
* 思路 DP
通过数组记录每次buy & sell & cooldown 后的最大值
buy[i] = Max(buy[i-1], cooldown[i-1]-prices[i])
sell[i] = Max(buy[i-1]+prices[i], sell[i-1])
cooldown[i] = sell[i-1]
