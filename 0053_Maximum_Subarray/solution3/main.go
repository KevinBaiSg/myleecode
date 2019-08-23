package solution3

func maxSubArray(nums []int) int {
	if len(nums) == 1 { return nums[0] }
	return divide(nums, 0, len(nums)-1)
}

func divide(nums []int, left, right int) int {
	const IntMax = int(^uint(0) >> 1)
	const IntMin = ^IntMax

	if len(nums) == 0 { return IntMin }

	if left > len(nums) - 1 { return IntMin}

	if left > right { return nums[left] } // 特殊情况，由于mid计算造成

	mid := left + (right - left) / 2

	leftMax := divide(nums, left, mid-1)
	rightMax := divide(nums, mid+1, right)
	midMax := findMidMax(nums, left, right, mid)

	return max(leftMax, rightMax, midMax)
}

func findMidMax(nums []int, left, right, mid int) int {
	const IntMax = int(^uint(0) >> 1)
	const IntMin = ^IntMax

	midLeftMax := IntMin
	if mid > 0 {
		midSum := nums[mid-1]
		midLeftMax = midSum
		for i := mid-2; i >= 0; i-- {
			midSum += nums[i]
			midLeftMax = max(midLeftMax, midSum)
		}
	}

	midRightMax := IntMin
	if mid+1 < right - left {
		midSum := nums[mid+1]
		midRightMax = midSum
		for i := mid+2; i <= right - left; i++ {
			midSum += nums[i]
			midRightMax = max(midRightMax, midSum)
		}
	}

	midMax := nums[mid]
	if midLeftMax != IntMin {
		midMax = max(midMax, midMax+midLeftMax)
	}
	if midRightMax != IntMin {
		midMax = max(midMax, midMax+midRightMax)
	}

	return midMax
}

func max(m ...int) int {
	if len(m) == 0 { return 0 }
	max := m[0]
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}
