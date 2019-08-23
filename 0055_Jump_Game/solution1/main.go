package solution1

func canJump(nums []int) bool {
	if len(nums) == 0 { return false }

	dp := make([]bool, len(nums))
	dp[len(nums)-1] = true

	for i := len(nums) - 2; i >= 0; i-- {
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			dp[i] = dp[i] || dp[i+j]
			if dp[i] { break }
		}
	}

	return dp[0]
}
