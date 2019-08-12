package sort

func HeapSort(array []int) []int {
	m := len(array)
	s := m / 2
	// 构建最大堆
	for i := s; i > -1; i-- {
		heap(array, i, m - 1)
	}
	// 堆排序
	for i := m - 1; i > 0; i-- {
		array[i], array[0] = array[0], array[i] // 移除根节点
		heap(array, 0, i - 1)
	}
	return array
}

func heap(array []int, i, end int) {
	l := 2 * i + 1 // 左子树节点
	if l > end { return }

	n := l // 标识左右节点大的一个，先标识为左节点
	r := 2 * i + 2	// 右子树节点
	if r <= end && array[r] > array[l] {
		n = r	// 右节点大于左节点，n 标识右节点
	}

	if array[i] > array[n] { return } // 不需要交换

	array[n], array[i] = array[i], array[n] // 交换父节点和大的子节点
	heap(array, n, end)	// 修整堆
}
