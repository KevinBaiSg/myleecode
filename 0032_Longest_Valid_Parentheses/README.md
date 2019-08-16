
### 问题描述
32. Longest Valid Parentheses
Hard

Given a string containing just the characters '(' and ')', 
find the length of the longest valid (well-formed) parentheses substring.

Example 1:
``` 
Input: "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()"
``` 

Example 2:
``` 
Input: ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()"
``` 

### 解决方法
#### DP
思路：
1. 状态定义：DP[i]: i 下标下最长有效 Parentheses
2. 状态转移方程：
    if s[i] == ")" && s[i-1] == "(" 
        DP[i] = DP[i-2] + 2
    if s[i] == ")" && s[i-1] == ")"
        if s[i-dp[i-1]-1] == "(" 
            dp[i] = dp[i-1] + 2 + dp[i-dp[i-1] - 2] 

如果 s[i] == ")" && s[i-1] == "("，也就是此时 `i-1` 和 `i` 是配对的，那么最长有效 Parentheses 就是 dp[i-2] 加上 2（s[i-1] 和 s[i]）
如果 s[i] == ")" && s[i-1] == ")"，这时候从字符串前边早可匹配的了，假设 dp[i-1] 匹配，那么 s[i-dp[i-1]-1] 位置为 "(" 也就可以匹配。
    这时可以通过 dp[i-1] + 2 计算可得。但是还要注意一种情况，比如 `()(())`，这种情况还需要加上之前最长有效 Parentheses，最后的这种情况的dp方程为
    dp[i] = dp[i-1] + 2 + dp[i-dp[i-1] - 2]
    
#### solution2
思路：Stack
解决 Parentheses 常用的方法

#### solution3
思路：官方的计数法
