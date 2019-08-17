package solution1

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 { return []int{-1, -1} }
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}

	left, right := 0, len(nums)
	targetL, targetR := -1, -1

	// binary search left
	left, right = 0, len(nums)
	for left < right {
		mid := left + (right - left) / 2
		if target == nums[mid] {
			right = mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left < len(nums) && nums[left] == target {
		targetL = left
	}

	// binary search right
	left, right = 0, len(nums)
	for left < right {
		mid := left + (right - left) / 2
		if target == nums[mid] {
			left = mid + 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left > 0 && nums[left-1] == target {
		targetR = left - 1
	}

	return []int{targetL, targetR}
}

