package solution1

func isPalindrome(s string) bool {
	if len(s) < 2 { return true }
	i, j := 0, len(s)-1

	for i < j {
		if !((s[i] >= '0' && s[i] <= '9') || (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z')) {
			i++; continue
		}
		if !((s[j] >= '0' && s[j] <= '9') || (s[j] >= 'a' && s[j] <= 'z') || (s[j] >= 'A' && s[j] <= 'Z')) {
			j--; continue
		}

		if s[i] == s[j] { // 包括数字、相同字母和特殊字符
			i++; j--; continue
		}

		// 不应该再存在数字的情况
		if (s[i] >= '0' && s[i] <= '9') || (s[j] >= '0' && s[j] <= '9') {
			return false
		}

		if (s[i] > s[j]) && (s[i] - s[j] == 'a' - 'A') {
			i++; j--; continue
		}

		if (s[i] < s[j]) && (s[j] - s[i] == 'a' - 'A') {
			i++; j--; continue
		}

		return false
	}

	return true
}
