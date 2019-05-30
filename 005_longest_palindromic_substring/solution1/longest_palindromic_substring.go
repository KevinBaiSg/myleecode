package solution1

func longestPalindrome(s string) string {
	r := []rune(s)
	if len(r) == 0 { return "" }
	if len(r) == 1 { return string(r) }
	/*
	DP:dp[i][j]: 回文开始和结束 j,i
	状态转移：if(dp[i+1][j-1] && s[i-1] == s[j+1]) {
				dp[i][j]=1
			}
	*/
	dp := make([][]int, len(r), len(r))
	for i := 0; i < len(r); i++ {
		dp[i] = make([]int, len(r), len(r))
	}

	start := 0
	end := 0

	for i := 0; i < len(r); i++ {
		dp[i][i] = 1
	}

	for i := 1; i < len(r) ; i++ {
		for j := 0; j < i ; j++ {
			if s[j] == s[i] && (dp[i-1][j+1] == 1 || i-1 < j+1) {
				dp[i][j] = 1
			}
			if dp[i][j] == 1 && i - j > end - start {
				start, end = j, i
			}
		}
	}
	return string(r[start:end+1])
}
