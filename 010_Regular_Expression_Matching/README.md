
### 问题描述


### 解决方法
思路：DP
1.状态定义：DP[i,j] 表示 s[i:] 是否匹配 p[j:]
2.状态方程：
    if p[j]==s[i] || p[j]=='.'
        DP[i,j]=DP[i+1,j+1]
    else if p[j]=='*'
        if p[j-1]==s[i] //一个
            DP[i,j]=DP[i-1,j-1] || DP[]
        DP[i,j]=
    for i:=len(s);i>0;i--
        for j:=len(p):j>0;j--
             
/*
		DP: DP[i,j] 表示 s[i:] 是否匹配 p[j:]
	*/