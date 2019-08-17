package common

/*
nums 升序
*/
func BinarySearch(nums []int, target int) int {
	if len(nums) == 0 { return -1 }
	if len(nums) == 1 {
		if nums[0] == target { return 0 }
		return -1
	}

	left, right := 0, len(nums) - 1
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
