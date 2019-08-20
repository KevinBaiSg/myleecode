package solution1

func rotate(matrix [][]int)  {
	width := len(matrix)

	if width < 2 { return }

	// 转置
	for i := 0; i < width; i++ {
		for j := i; j < width; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 翻转
	for i := 0; i < width; i++ {
		for j := 0; j < width/2; j++ {
			matrix[i][j], matrix[i][width-j-1] = matrix[i][width-j-1], matrix[i][j]
		}
	}
}
