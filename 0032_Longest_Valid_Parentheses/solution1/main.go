package solution1

func longestValidParentheses(s string) int {
	/*
	   DP思路
	   1.DP[i] 表示 i 标识字符串下 最长有效 parentheses
	   2. if s[i] == ")" && s[i-1] == "("
	           DP[i] = DP[i-2] + 2
	       if s[i] == ")" && s[i-1] == ")"
	           if s[i-dp[i-1]-1] == "("
	               dp[i] = dp[i-1] + 2 + dp[i-dp[i-1] - 2]
	*/

	if len(s) < 2 { return 0 }

	curMax := 0
	dp := make([]int, len(s), len(s))

	for i := 1; i < len(s); i++ {
		if s[i] == ')' && s[i-1] == '(' {
			if i >= 2 {
				dp[i] =  dp[i - 2] + 2
			} else {
				dp[i] = 2
			}
		} else if s[i] == ')' && s[i - 1] == ')' {
			if i - dp[i - 1] > 0 && s[i - dp[i - 1] - 1] == '(' {
				if i - dp[i - 1] >= 2 {
					dp[i] = dp[i - 1] + 2 + dp[i - dp[i - 1] - 2]
				} else {
					dp[i] = dp[i - 1] + 2
				}
			}
		}

		curMax = max(curMax, dp[i])
	}

	return curMax
}

func max(a, b int) int {
	if a < b { return b }
	return a
}
