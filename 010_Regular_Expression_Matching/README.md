
### 问题描述


### 解决方法
#### solution1
/*
		DP: DP[i,j] 表示 s[i:] 是否匹配 p[j:]

		p[j] == s[i]:dp[i][j] = dp[i-1][j-1]
		p[j] == ".":dp[i][j] = dp[i-1][j-1]
		p[j] =="*":
			3.1 p[j-1] != s[i]:dp[i][j] = dp[i][j-2]
			3.2 p[i-1] == s[i] or p[i-1] == ".":
			dp[i][j] = dp[i-1][j] // 多个a的情况
			or dp[i][j] = dp[i][j-1] // 单个a的情况
			or dp[i][j] = dp[i][j-2] // 没有a的情况

			转移方程：
				s: abddcecf	无
				p: abddcc*ecf
				=========
				s: abcef	一个
				p: abc*ef
				==============
				s: abccccef 多个
				p: abc*ef

				for i:=0; i<len(s);i++
					for j=0;j<len(p);j++
						if p[j] == "*"
							if s[i]==p[j-1] || p[j-1]=="."
								DP[i,j]=DP[i-2,j-2] || DP[i-1,j-1] || DP[i-1,j]	//
							else
								DP[i,j]=DP[i,j-2] // "*" 为 0
						else if p[j]=="."
							DP[i,j]=DP[i-1,j-1]
						else
							DP[i,j]=DP[i-1,j-1] && s[i] == p[j]
	*/