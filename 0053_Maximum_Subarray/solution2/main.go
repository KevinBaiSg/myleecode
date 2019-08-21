package solution1

func maxSubArray(nums []int) int {
	if len(nums) == 1 { return nums[0] }

	dp := make([]int, len(nums)) // 压缩状态，其实就跟 solution1 相同了
	dp[0] = nums[0]
	curMax := dp[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		curMax = max(curMax, dp[i])
	}

	return curMax
}

func max(a, b int) int {
	if a < b { return b }
	return a
}
