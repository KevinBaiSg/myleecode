package solution3

func topK(nums []int, k int) []int {
	if len(nums) <= k { return nums }

	heapTopK := NewHeapTopK(nums, k)
	heapTopK.MakeMinHeap()

	for i := k; i < len(nums); i++ {
		if nums[i] > heapTopK.Min() {
			nums[i] = heapTopK.ReplaceMin(nums[i])
			heapTopK.AdjustHeap(0)
		}
	}
	return nums[:k]
}

type HeapTopK struct {
	nums 	[]int
	k 		int
}

func NewHeapTopK(nums []int, k int) *HeapTopK {
	if k > len(nums) { return nil }

	ht := HeapTopK{
		nums:	nums,
		k:		k,
	}
	return &ht
}

func (h *HeapTopK)MakeMinHeap() {
	m := h.k / 2
	for i := m; i >= 0; i-- {
		h.AdjustHeap(i)
	}
}

func (h *HeapTopK)AdjustHeap(i int) {
	l := 2 * i + 1  // 左节点 index
	r := 2 * i + 2	// 右节点 index

	if l > h.k { return }

	min := l
	if r < h.k && h.nums[r] < h.nums[l] {
		//h.nums[l], h.nums[r] = h.nums[r], h.nums[l] // 交换左右节点
		min = r
	}

	if h.nums[i] > h.nums[min] {
		h.nums[i], h.nums[min] = h.nums[min], h.nums[i] // 交换节点
	}

	h.AdjustHeap(min)
}

func (h *HeapTopK)Min() int {
	return h.nums[0]
}

func (h *HeapTopK)ReplaceMin(new int) (old int) {
	old, h.nums[0] = h.nums[0], new
	return
}
