package _443_String_Compression

import "strconv"

func compress(chars []byte) int {
	var (
		current = chars[0]
		index = 0	// 已设置的位置
		count = 1	// 重复次数
	)

	for i := 1; i < len(chars); i++ {
		if chars[i] == current {
			count ++
		} else {
			// todo copy count
			if count > 1 {
				countStr := strconv.Itoa(count)
				for i := 0; i < len(countStr) ; i++ {
					index ++
					chars[index] = countStr[i]
				}
			}

			// reset
			count = 1
			current = chars[i]
			index ++
			chars[index] = current
		}
	}

	if count > 1 {
		countStr := strconv.Itoa(count)
		for i := 0; i < len(countStr) ; i++ {
			index ++
			chars[index] = countStr[i]
		}
	}

	return index + 1
}