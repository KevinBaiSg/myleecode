package solution1

func maxArea(height []int) int {
	if len(height) < 2 { return 0 }

	l, r := 0, len(height) - 1
	maxAre := 0

	for l < r {
		if height[l] > height[r] {
			maxAre = Max(maxAre, (r - l) * height[r])
			r--
		} else {
			maxAre = Max(maxAre, (r - l) * height[l])
			l++
		}
	}

	return maxAre
}

func Max(a, b int) int {
	if a > b { return a }
	return b
}
