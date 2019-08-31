package solution1

func numDistinct(s string, t string) int {
	if len(t) > len(s) { return 0 }
	if s == t { return 1 }

	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
		dp[i][0] = 1
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(s)][len(t)]
}
