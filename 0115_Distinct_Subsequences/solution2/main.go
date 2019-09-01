package solution2

import "strconv"

func numDistinct(s string, t string) int {
	if len(t) > len(s) { return 0 }
	if s == t { return 1 }
	mp := make(map[string]int)

	return dfs(s, 0, t, 0, &mp, 0)
}

func dfs(s string, sIndex int, t string, tIndex int, mp *map[string]int, curCount int) (count int) {
	count = curCount

	if tIndex == len(t) {
		count++; return
	}

	if sIndex == len(s) { return }
	if len(s) - sIndex < len(t) - tIndex { return } // 剪枝

	key := strconv.Itoa(sIndex) + ":" + strconv.Itoa(tIndex)
	if value, ok := (*mp)[key]; ok { return value }

	if s[sIndex] == t[tIndex] {
		count = dfs(s, sIndex+1, t, tIndex+1, mp, count) + dfs(s, sIndex+1, t, tIndex, mp, count)
	} else {
		count = dfs(s, sIndex+1, t, tIndex, mp, count)
	}

	(*mp)[key] = count
	return
}
