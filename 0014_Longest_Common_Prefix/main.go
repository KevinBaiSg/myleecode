package _14_Longest_Common_Prefix

/*
	循环列查找
*/

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 { return "" }

	var i = 0
	
	for ;; {
		var current byte = 0x00
		for j, str := range strs {
			if i < len(str) {
				if j == 0 {
					current = str[i]
				} else {
					if current != str[i] {
						return str[0:i]
					}
				}
			} else {
				return str
			}
		}
		i ++
	}
}
