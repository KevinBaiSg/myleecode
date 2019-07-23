package solution1

func isScramble(s1 string, s2 string) bool {
	// 边界判断
	if s1 == s2 { return true }
	if len(s2) != len(s2) { return false }

	// 字符匹配判断
	n := len(s1)
	var counts = make([]int, 26, 26)
	for i := 0; i < n; i++ {
		counts[s1[i] - 'a'] ++
		counts[s2[i] - 'a'] --
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	// for range s 处理
	for i := 1; i < n; i++ {
		// 直接对比两子树
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}

		// 交换两子树，然后对比
		if isScramble(s1[i:], s2[:n-i]) && isScramble(s1[:i], s2[n-i:]) {
			return true
		}
	}

	return false
}