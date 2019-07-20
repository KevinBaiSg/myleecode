package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	runeMap := make(map[rune]int)
	longest := 0
	head, end := -1, -1

	for i, rune := range runes {
		point, ok := runeMap[rune]
		if ok && point > head {
			head = point
		}
		end = i
		if longest < end-head {
			longest = end - head
		}
		runeMap[rune] = i
	}
	return longest
}
