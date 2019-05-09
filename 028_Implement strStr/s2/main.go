package s2

/*
ref:
https://www.youtube.com/watch?v=dgPabAsTFa8
https://www.youtube.com/watch?v=3IFxpozBs2I
*/

func getPrefix(needle string) []int {
	prefix := make([]int, len(needle))
	var (
		i = 0
		index = -1
	)
	/*
		if index == -1
			prefix[0] = -1
		if needle[i] == needle[index]
			prefix[i] = index + 1
		else

	*/
	prefix[0] = -1
	for i < len(needle) - 1 {
		if index == -1 || needle[i] == needle[index] {
			index ++
			i ++
			prefix[i] = index
		} else {
			index = prefix[index] // 向前比较
		}
	}

	return prefix
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 { return 0 }
	prefix := getPrefix(needle)

	var (
		i = 0	// index for haystack
		j = 0	// index for needle
	)
	for i < len(haystack) && j < len(needle)  {
		if j == -1 || haystack[i] == needle[j] {
			i ++
			j ++
		} else {
			j = prefix[j]
		}
	}

	if j == len(needle) {
		return i - j
	}
	return -1
}