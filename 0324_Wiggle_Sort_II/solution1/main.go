package solution1

func wiggleSort(nums []int)  {
	/*
	   1. quickSelect(快速选择算法)求得中间数
	   2. three-way-partition 方法交换对于位置
	*/

	// get top k
	k := (len(nums) + 1) / 2
	// quick select
	mid := quickSelect(nums, 0, len(nums)-1, k)
	// three way partition
	threeWayPartition(nums, mid)
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

func threeWayPartition(nums []int, mid int)  {
	n, i, l, r := len(nums), 0, 0, len(nums)-1
	for i <= r {
		iM, lM, rM := indexMap(i, n), indexMap(l, n), indexMap(r, n)
		if nums[iM] > mid {
			nums[lM], nums[iM] = nums[iM], nums[lM]
			l++; i++
		} else if nums[iM] < mid {
			nums[rM], nums[iM] = nums[iM], nums[rM]
			r--
		} else {
			i++
		}
	}
}

func indexMap(i, n int) int {
	return (2 * i + 1) % (n | 1)
}
