package solution1

func numDistinct(s string, t string) int {
	if len(t) > len(s) { return 0 }
	if s == t { return 1 }

	dp := make([]int, len(t)+1)
	dp[0] = 1

	for i := 1; i <= len(s); i++ {
		for j := len(t); j > 0; j-- {
			if s[i-1] == t[j-1] {
				dp[j] += dp[j-1]
			}
		}
	}

	return dp[len(t)]
}

