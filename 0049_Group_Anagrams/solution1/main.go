package solution1

import "sort"

func groupAnagrams(strs []string) [][]string {
	keyMap := make(map[string]int)
	rets := make([][]string, 0)

	for _, str := range strs {
		key := runeSort(str)
		sort.Sort(key)
		if v, ok := keyMap[string(key)]; ok {
			rets[v] = append(rets[v], str)
		} else {
			rets = append(rets, []string{ str })
			keyMap[string(key)] = len(rets) - 1
		}
	}
	return rets
}

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
