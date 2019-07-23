
### 问题描述
```
87. Scramble String
Hard

Given a string s1, we may represent it as a binary tree by partitioning it to two non-empty substrings recursively.

Below is one possible representation of s1 = "great":

    great
   /    \
  gr    eat
 / \    /  \
g   r  e   at
           / \
          a   t
To scramble the string, we may choose any non-leaf node and swap its two children.

For example, if we choose the node "gr" and swap its two children, it produces a scrambled string "rgeat".

    rgeat
   /    \
  rg    eat
 / \    /  \
r   g  e   at
           / \
          a   t
We say that "rgeat" is a scrambled string of "great".

Similarly, if we continue to swap the children of nodes "eat" and "at", it produces a scrambled string "rgtae".

    rgtae
   /    \
  rg    tae
 / \    /  \
r   g  ta  e
       / \
      t   a
We say that "rgtae" is a scrambled string of "great".

Given two strings s1 and s2 of the same length, determine if s2 is a scrambled string of s1.

Example 1:

Input: s1 = "great", s2 = "rgeat"
Output: true
Example 2:

Input: s1 = "abcde", s2 = "caebd"
Output: false

```  

### 解决方法
#### solution1
思路：递归
如果子树符合问题，那么总体就符合问题，并且子树的解法与原问题想同，所以可以使用递归来接此题。
并且在正式求解之前可以使用一些简单条件判断排除一些问题。


#### solution2
思路：DP
1.状态定义: dp[i] 
dp[i][j]: 回文开始和结束 j,i
2.状态转移: if s[j] == s[i] && (dp[i-1][j+1] == 1 || i-1 < j+1) {
            dp[i][j] = 1
            }
