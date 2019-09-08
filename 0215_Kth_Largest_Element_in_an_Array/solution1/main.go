package solution1

func findKthLargest(nums []int, k int) int {
	if k == 1 && len(nums) == 1 { return nums[0] }
	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, first, end int, k int) int {
	if first == end { return nums[first] }

	i := partition(nums, first, end)
	if i-first+1 == k {
		return nums[i]
	} else if i-first+1 > k {
		return quickSelect(nums, first, i-1, k)
	} else {
		return quickSelect(nums, i+1, end, k-(i-first+1))
	}
}

func partition(nums []int, first, end int) int {
	l, r := first, end
	i := l
	pivot := nums[l]
	for i <= r {
		if nums[i] > pivot {
			nums[l], nums[i] = nums[i], nums[l]
			i++; l++
		} else if nums[i] < pivot {
			nums[r], nums[i] = nums[i], nums[r]
			r--
		} else {
			i++
		}
	}

	return i-1
}