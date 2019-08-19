package solution1

func trap(height []int) int {
	if len(height) < 2 { return 0 }

	left, right := 0, len(height) - 1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		if height[left] < height[right] { // left
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else { // right
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}

	return water
}