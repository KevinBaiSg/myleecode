
### 问题描述
22. Generate Parentheses
Medium

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

```     
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

``` 

### 解决方法
#### solution1
思路：回溯方法
深度搜索组合，但是要符合一定的规则：
1.left >= right 
2. left right 都小于 n