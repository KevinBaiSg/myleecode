package sort

func QuickSort(array []int) []int {
	if len(array) < 2 { return array }

	pivot := array[0]
	l, r := 0, len(array) - 1

	for i := 1; i <= r; {
		if array[i] < pivot { //
			array[l], array[i] = array[i], array[l]
			l++; i++
		} else {
			array[r], array[i] = array[i], array[r]
			r--
		}
	}

	QuickSort(array[:l])
	QuickSort(array[l + 1:])
	return array
}
