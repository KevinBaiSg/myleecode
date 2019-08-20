package solution2

func trap(height []int) int {
	if len(height) < 2 { return 0 }

	dpLeft := make([]int, len(height), len(height))
	dpRight := make([]int, len(height), len(height))
	water := 0

	for i := 1; i < len(height); i++ {
		dpLeft[i] = max(dpLeft[i-1], height[i-1])
	}

	for i := len(height) - 2; i >= 0; i-- {
		dpRight[i] = max(dpRight[i+1], height[i+1])
	}

	for i := 0; i < len(height); i++ {
		lower := min(dpLeft[i], dpRight[i])
		if lower > height[i] {
			water += lower - height[i]
		}
	}

	return water
}

func max(a, b int) int {
	if a < b { return b }
	return a
}

func min(a, b int) int {
	if a > b { return b }
	return a
}