
### 问题描述
```
122. Best Time to Buy and Sell Stock II
Easy

Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit.
You may complete as many transactions as you like
(i.e., buy one and sell one share of the stock multiple times).

Note: You may not engage in multiple transactions at the same time
(i.e., you must sell the stock before you buy again).

Example 1:

Input: [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
             Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Example 2:

Input: [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
             Note that you cannot buy on day 1, buy on day 2 and sell them later,
            as you are
             engaging multiple transactions at the same time.
            You must sell before buying again.
Example 3:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
```  

### 解决方法
#### Solution 1
* 思路  
数组图像中最近差值等于最高点与最低点差值，或者即使有峰谷的数组中，可以把所有增长曲线连接起来，即为最优值

#### Solution 2
* 思路 DP
MaxProfit=MP 最大收益
k 最大交易次数
1.状态定义：MP[i,j]   到i天最大收益
                    j=0:没有股票 j=1:拥有一股
                             
2.  MP[i,0]=Max(MP[i-1,0],MP[i-1,1]+a[i])
    MP[i,1]=Max(MP[i-1,0]-a[i],MP[i-1,1])
