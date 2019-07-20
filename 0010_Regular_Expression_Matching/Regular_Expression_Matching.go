//package _10_Regular_Expression_Matching
package main

func main()  {
	isMatch("aab", "c*a*b")
}
func isMatch(s string, p string) bool {
	if s == p { return true }

	dp := make([][]bool, len(s)+1, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1, len(p)+1)
	}

	dp[0][0] = true
	for j := 1; j < len(p); j++ {
		if p[j] == '*' && dp[0][j-1] {
			dp[0][j+1] = true
		}
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if p[j] == s[i] || p[j] == '.' {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				if p[j - 1] != s[i] {
					dp[i+1][j+1] = dp[i+1][j-1]
				}
				if s[i] == p[j-1] || p[j-1] == '.' {
					dp[i+1][j+1] = dp[i][j+1] || dp[i+1][j] || dp[i+1][j-1]
				}
			}
		}
	}

	return dp[len(s)][len(p)]
}