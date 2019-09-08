package solution4

func topK(nums []int, k int) []int {
	if len(nums) <= k { return nums }
	quickTop(nums, 0, len(nums) - 1, k)
	return nums[:k]
}

func quickTop(nums []int, first, end, k int) {
	if first == end { return }

	i := partition(nums, first, end)

	if i-first+1 == k {
		return
	} else if i-first+1 > k {
		quickTop(nums, first, i-1, k)
	} else {
		quickTop(nums, i+1, end, k-(i-first+1))
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
