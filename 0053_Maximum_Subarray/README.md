
### 问题描述
53. Maximum Subarray
Easy

Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

```text
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```

Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

### 解决方法
#### solution1
思路：Kadane算法
这应该是该问题的最优解了，参照[维基百科](https://zh.wikipedia.org/wiki/%E6%9C%80%E5%A4%A7%E5%AD%90%E6%95%B0%E5%88%97%E9%97%AE%E9%A2%98)


#### solution2
思路：DP
DP 状态定义 dp[i]: 到 i 位置是的最大和，但是 nums[i] 必须包括在内
DP 状态转移方程:
```go
if dp[i-1] > 0 {
	dp[i]= dp[i-1] + nums[i]
} else {
	dp[i] = nums[i]
}
```
这时候我们可以进一步优化，做状态压缩，其实跟solution1就一样了