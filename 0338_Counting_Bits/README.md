
### 问题描述
338. Counting Bits
Medium

Given a non negative integer number num. For every numbers i in the range 0 ≤ i ≤ num calculate the number of 1's in their binary representation and return them as an array.

Example 1:

```text
Input: 2
Output: [0,1,1]
```

Example 2:

```text
Input: 5
Output: [0,1,1,2,1,2]
```

Follow up:

It is very easy to come up with a solution with run time O(n*sizeof(integer)). But can you do it in linear time O(n) /possibly in a single pass?
Space complexity should be O(n).
Can you do it like a boss? Do it without using any builtin function like __builtin_popcount in c++ or in any other language.

### 解决方法
#### solution1
思路：DP, x & (x-1)
我们还可以借助 `x & (x-1)` 消除最后一个 `1` 这个特质，只不过这次是在下标上。
定义 `dp[i]` 为 `i` 位置的 `1` 的数量，定义 `dp[i] = dp[i & (i-1)] + 1`。这个公司的意义就是 `i & (i-1)` 与 `i` 相差最后一个 `1`。

#### solution2
思路：寻找 1 数量和 i 的关系