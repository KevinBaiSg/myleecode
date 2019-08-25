
### 问题描述
509. Fibonacci Number
Easy

The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), for N > 1.
Given N, calculate F(N).

Example 1:

```text
Input: 2
Output: 1
Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
```

Example 2:

```text
Input: 3
Output: 2
Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
```

Example 3:

```text
Input: 4
Output: 3
Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.
```
 
Note:
0 ≤ N ≤ 30.

### 解决方法
#### solution1
思路：递归
但是递归的时间复杂度为指数级的增长

#### solution2
思路：正推法

从斐波那契数列的定义：
f(0) = 0;
f(1) = 1;
f(n) = f(n-1) + f(n-2) n>=2时

其实我们正向计算 f(0),f(1),f(2)…f(n)，最后实现的代码我们可以做到时间复杂度O(n)，空间复杂度O(1)

### 参考
[拜托，面试别再问我斐波那契数列了！！！](https://mp.weixin.qq.com/s/3LR-iVC4zgj0tGhZ780PcQ)
