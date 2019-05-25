package _20_Triangle

/*
	120. Triangle
	Medium

	Given a triangle, find the minimum path sum from top to bottom.
	Each step you may move to adjacent numbers on the row below.

	For example, given the following triangle

	[
		 [2],
		[3,4],
	   [6,5,7],
	  [4,1,8,3]
	]
	The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

	Note:
	Bonus point if you are able to do this using only O(n) extra space,
	where n is the total number of rows in the triangle.
*/

/*
	解题思路：动态规划
	1.DP 状态定义 DP[i,j]: bottom->(i,j) path sum min
	2.DP 状态方程：DP[i,j]=min(DP[i+1,j], DP[i+1,j+1]) + Triangle[i,j]
*/
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	//mini := make([]int, m, m) // 只记录下层，状态压缩

	/*
		初始化最底层最小状态，DP[m-1][j] = triangle[m - 1][j]
	*/
	// copy(mini, triangle[m - 1])
	// or 借用 triangle[m - 1]
	mini := triangle[m - 1]

	// 递推向上求解最小状态
	for i := m - 2; i >= 0; i--  {
		for j := 0; j < len(triangle[i]); j++  {
			if mini[j] < mini[j + 1]{
				mini[j] = triangle[i][j] + mini[j]
			} else {
				mini[j] = triangle[i][j] + mini[j + 1]
			}
		}
	}

	return mini[0]
}