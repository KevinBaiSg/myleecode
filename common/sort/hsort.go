package sort

func HeapSort(array []int) []int {
	m := len(array)
	s := m/2
	// 构建堆
	for i := s; i > -1; i-- {
		heap(array, i, m-1)
	}
	for i := m-1; i > 0; i-- {
		array[i], array[0] = array[0], array[i]
		heap(array, 0, i-1)
	}
	return array
}

func heap(array []int, i, end int) {
	l := 2 * i + 1 // 左子树节点
	if l > end { return }

	n := l
	r := 2 * i + 2	// 右子树节点
	if r <= end && array[r] > array[l] {
		n = r
	}

	if array[i] > array[n]{ return }

	array[n], array[i] = array[i], array[n]
	heap(array, n, end)
}
