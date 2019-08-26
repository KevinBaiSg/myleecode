
### 问题描述
231. Power of Two
Easy

Given an integer, write a function to determine if it is a power of two.

Example 1:

```text
Input: 1
Output: true 
Explanation: 20 = 1
```

Example 2:

```text
Input: 16
Output: true
Explanation: 24 = 16
```

Example 3:

```text
Input: 218
Output: false
```

### 解决方法
#### solution1
思路：n & (n-1) 来判断是否存在 `1`，因为 2 的 n 次幂只有一个 `1`
