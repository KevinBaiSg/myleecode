package Roman_to_Integer

func romanToInt(s string) int {
	romanMap := map[rune]int{
		rune('I'): 1,
		rune('V'): 5,
		rune('X'): 10,
		rune('L'): 50,
		rune('C'): 100,
		rune('D'): 500,
		rune('M'): 1000,
	}

	runes := []rune(s)
	if len(runes) == 1 {
		val, ok := romanMap[runes[0]]
		if !ok {
			return 0
		}
		return val
	}

	number := 0
	for i := 0; i < len(runes); i++ {
		next := i + 1
		if next < len(runes) {
			if romanMap[runes[i]] < romanMap[runes[next]] {
				number += romanMap[runes[next]] - romanMap[runes[i]]
				i++
				continue
			}
		}
		number += romanMap[runes[i]]
	}

	return number
}
