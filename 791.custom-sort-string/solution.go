package leetcode

import (
	"sort"
)

var (
	strMap = make(map[byte]int, 0)
)

func customSortString(S string, T string) string {
	for i, v := range S {
		strMap[byte(v)] = i
	}
	tmp := CustomSort(T)
	sort.Sort(tmp)
	return string(tmp)

}

type CustomSort []byte

func (s CustomSort) Len() int {
	return len(s)
}

func (s CustomSort) Less(i int, j int) bool {
	return strMap[s[i]] < strMap[s[j]]
}

func (s CustomSort) Swap(i int, j int) {
	tmp := []byte(s)
	tmp[i], tmp[j] = tmp[j], tmp[i]
}
