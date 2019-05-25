package _52_Maximum_Product_Subarray
/*
	152. Maximum Product Subarray
	Medium

	Given an integer array nums, find the contiguous subarray within
	an array (containing at least one number) which has the largest product.

	Example 1:

	Input: [2,3,-2,4]
	Output: 6
	Explanation: [2,3] has the largest product 6.
	Example 2:

	Input: [-2,0,-1]
	Output: 0
	Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
*/

/*
	解题思路：DP
	1.状态定义：DP[i][2]: 0 -> i: max and min(负)
	2.DP转移方程：
		if (a[i] >= 0) {
			DP[i][0]=DP[i-1][0]*a[i]
			DP[i][1]=DP[i-1][1]*a[i]
		} else {
			DP[i][0]=DP[i-1][1]*a[i]
			DP[i][1]=DP[i-1][1]*a[i]
		}

*/
func maxProduct(nums []int) int {
	// 边界
	if nums == nil || len(nums) == 0 { return 0 }
	if len(nums) == 1 { return nums[0] }

	// 初始化
	res, curMax, curMin := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++  {
		curMax, curMin = MaxMin(curMax*nums[i], curMin*nums[i], nums[i])
		res = Max(curMax, res)
	}

	return res
}

func Max(m ...int) int {
	if len(m) == 0 { return 0 }
	max := m[0]
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}

func MaxMin(m ...int) (int, int) {
	if len(m) == 0 { return 0, 0 }
	max, min := m[0], m[0]
	for _, v := range m {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max, min
}
