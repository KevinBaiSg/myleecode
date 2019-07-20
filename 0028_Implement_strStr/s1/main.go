package s1

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	var (
		i = 0
		j = 0
	)
	for i = 0; i < len(haystack); i++ {
		for j = 0; j < len(needle); j++  {
			if i + j < len(haystack) {
				if haystack[i + j] != needle[j] {
					break
				} else {
					if j == len(needle) - 1 {
						return i
					}
				}
			} else {
				return -1
			}
		}
	}
	return -1
}