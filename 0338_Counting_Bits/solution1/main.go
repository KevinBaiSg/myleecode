package solution1

func countBits(num int) []int {
	dp := make([]int, num+1, num+1) // 0

	for i := 1; i <= num; i++ {
		dp[i] = dp[i & (i-1)] + 1
	}

	return dp
}
