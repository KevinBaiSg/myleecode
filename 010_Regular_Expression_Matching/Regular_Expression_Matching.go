//package _10_Regular_Expression_Matching
package main

func main()  {
	isMatch("aab", "c*a*b")
}
func isMatch(s string, p string) bool {
	if s == p { return true }
	if len(s) == 0 || len(p) == 0 { return false }

	dp := make([][]bool, len(s), len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(p), len(p))
	}
	
	if p[0] == s[0] || p[0] == '.' {
		dp[0][0] = true
	}

	for j := 1; j < len(p); j++ {
		if p[j] == '*' && dp[0][j-1] {
			dp[0][j] = true
		}
	}

	for i := 1; i < len(s); i++ {
		for j := 1; j < len(p); j++ {
			if p[j] == '*' {
				if i > 1 && j > 1 {
					dp[i][j] = dp[i-2][j-2] || dp[i-1][j-1] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i-1][j-1] || dp[i-1][j]
				}
			} else if s[i] == p[j-1] || p[j] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] && s[i] == p[j]
			}
		}
	}

	return dp[len(s) - 1][len(p) - 1]
}