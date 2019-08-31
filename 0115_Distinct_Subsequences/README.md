
### 问题描述
115. Distinct Subsequences
Hard

Given a string S and a string T, count the number of distinct subsequences of S which equals T.

A subsequence of a string is a new string which is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (ie, "ACE" is a subsequence of "ABCDE" while "AEC" is not).

Example 1:

```text
Input: S = "rabbbit", T = "rabbit"
Output: 3
Explanation:

As shown below, there are 3 ways you can generate "rabbit" from S.
(The caret symbol ^ means the chosen letters)

rabbbit
^^^^ ^^
rabbbit
^^ ^^^^
rabbbit
^^^ ^^^
```

Example 2:

```text
Input: S = "babgbag", T = "bag"
Output: 5
Explanation:

As shown below, there are 5 ways you can generate "bag" from S.
(The caret symbol ^ means the chosen letters)

babgbag
^^ ^
babgbag
^^    ^
babgbag
^    ^^
babgbag
  ^  ^^
babgbag
    ^^^
```

### 解决方法
#### solution1
思路：DP
1.状态定义：dp[i][j] 表示 t 前 j 字符串可以由 s 前 i 字符串组成的最多 个数 
2.状态转移方程：
    if s[i] == t[j] {
        dp[i][j] = dp[i-1][j-1] + dp[i-1][j]    
    } else {
        dp[i][j] dp[i-1][j]
    }
    
dp[i-1][j-1]: 表明 s[i] 和 t[j] 相等，只要 dp[i-1][j-1] 匹配代表 dp[i][j] 也匹配 
dp[i-1][j]：表示 t[:j] 匹配 s[:i-1] 之前的值，所以需要加上

#### solution2
思路：递归+回溯

#### solution3
思路：递归+分治