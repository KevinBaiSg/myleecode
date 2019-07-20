
### 问题描述
```
5. Longest Palindromic Substring
Medium

Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
```  

### 解决方法

#### Solution1
* 思路 DP
1.状态定义: dp[i][j]: 回文开始和结束 j,i
2.状态转移: if s[j] == s[i] && (dp[i-1][j+1] == 1 || i-1 < j+1) {
            dp[i][j] = 1
            }
			
            