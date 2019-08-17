package solution1

func search(nums []int, target int) int {
	if len(nums) == 0 { return -1 }
	if len(nums) == 1 {
		if nums[0] != target { return -1 }
		return 0
	}

	rotateIndex := rotateIndexSearch(nums, 0, len(nums) - 1)
	if nums[rotateIndex] == target {
		return rotateIndex
	} else if rotateIndex == 0 {
		return binarySearch(nums, 0, len(nums) - 1, target)
	} else if rotateIndex == len(nums) - 1 {
		return binarySearch(nums, 0, len(nums) - 2, target)
	} else if nums[0] > target {
		return binarySearch(nums, rotateIndex + 1, len(nums) - 1, target)
	} else {
		return binarySearch(nums, 0, rotateIndex - 1, target)
	}
}

func rotateIndexSearch(nums []int, left, right int) int {
	if right - left == 0 { return 0 }

	for left <= right {
		mid := left + (right - left) / 2
		if mid == left {
			if nums[left] > nums[right] {
				return right
			} else {
				return 0
			}
		} else if nums[mid] > nums[mid+1] {
			return mid+1
		} else {
			if nums[mid] > nums[left] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}

	return 0
}

func binarySearch(nums []int, left, right int, target int) int {
	for left <= right {
		mid := left + (right - left) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
