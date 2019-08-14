package solution1

var letterMap = []string{
	" ",    //0
	"",     //1
	"abc",  //2
	"def",  //3
	"ghi",  //4
	"jkl",  //5
	"mno",  //6
	"pqrs", //7
	"tuv",  //8
	"wxyz", //9
}

var comLetters []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 { return []string{} }

	comLetters = []string{} // 初始化 comLetters
	backtrack(digits, 0, "")

	return comLetters
}

func backtrack(digits string, index int, s string) {
	letters := letterMap[digits[index] - '0']
	for i := 0; i < len(letters); i++ {
		if index == len(digits) - 1 {
			comLetters = append(comLetters, s + string(letters[i])) // 深度优先遍历最底层
		} else {
			backtrack(digits, index + 1, s + string(letters[i])) // 向下一层搜索
		}
	}
}
