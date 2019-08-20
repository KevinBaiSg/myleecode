
### 问题描述
42. Trapping Rain Water
Hard

Given n non-negative integers representing an elevation map where the width of each bar is 1, 
compute how much water it is able to trap after raining.

![RainWaterTrap](https://assets.leetcode.com/uploads/2018/10/22/rainwatertrap.png)

The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. 
In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!

Example:

```text
Input: [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
``` 


### 解决方法
#### solution1
思路：双指针

#### solution2
思路：DP
状态定义：dpLeft[i], dpRight[i]: 分别表示 i 位置左右两边最高墙
状态方程：dpLeft[i] = max(dpLeft[i-1], height[i-1])
        dpRight[i] = max(dpRight[i-1], height[i-1])
        计算时注意边界
        
        