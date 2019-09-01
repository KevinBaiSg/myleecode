
### 问题描述
477. Total Hamming Distance
Medium

The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Now your job is to find the total Hamming distance between all pairs of the given numbers.

Example:

```text
Input: 4, 14, 2

Output: 6

Explanation: In binary representation, the 4 is 0100, 14 is 1110, and 2 is 0010 (just
showing the four bits relevant in this case). So the answer will be:
HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2 + 2 + 2 = 6.
```

Note:
Elements of the given array are in the range of 0 to 10^9
Length of the array will not exceed 10^4.

### 解决方法
#### solution1
思路：一种比较巧妙的方法
统计每个二进制位 1 的个数和 0 的个数, 按照每个二进制位的不同将数组分为两组；一组是第 i 位为 1，一组是第 i 位为 0; 数据的两两组合表示为这两组数据之间的全排列 即其个数相乘

按位操作总共遍历32次，时间复杂度是 O(32 * n)=O(n)
