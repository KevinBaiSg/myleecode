package sort

import "sort"

type runeSort []rune

func (r runeSort) Len() int {
	return len(r)
}

func (r runeSort) Less(i, j int) bool {
	if r[i] < r[j] { return true }
	return false
}

func (r runeSort) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func InterfaceSort(src string) string {
	s := runeSort(src)
	sort.Sort(s)
	return string(s)
}