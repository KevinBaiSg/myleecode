package solution2

func canJump(nums []int) bool {
	if len(nums) == 0 { return false }

	lastPosition := len(nums)-1

	for i := len(nums)-2; i >= 0; i-- {
		if i+nums[i] >= lastPosition {
			lastPosition = i
		}
	}

	return lastPosition == 0
}
