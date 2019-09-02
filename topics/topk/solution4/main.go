package solution4

func topK(nums []int, k int) []int {
	if len(nums) <= k { return nums }
	quickTop(nums, 0, len(nums) - 1, k)
	return nums[:k]
}

func quickTop(nums []int, first, end, k int) {
	if first == end { return }

	pivot := nums[first]
	l, r := first, end
	for i := l+1; i < r; {
		if nums[i] > pivot {
			nums[l], nums[i] = nums[i], nums[l]
			l++; i++
		} else {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		}
	}
	if l - first + 1 == k {
		return
	} else if l - first > k {
		quickTop(nums, first, l - 1, k)
	} else {
		quickTop(nums, l, end, k+first-l)
	}
}
