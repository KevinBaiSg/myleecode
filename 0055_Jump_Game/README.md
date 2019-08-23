
### 问题描述
55. Jump Game
Medium

Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

Example 1:

```text
Input: [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
```

Example 2:

```text

Input: [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum
             jump length is 0, which makes it impossible to reach the last index.
```


### 解决方法
#### solution1
思路：DP
状态定义：DP[i]: nums[i] 能否到达结尾
状态转移方程：DP[i]=
            for j := 0; j <= nums[i]; j++
                {
                    DP[i] || DP[i+j] // i+j < len(nums)
                }
            
#### solution1
思路：xxxxxxx