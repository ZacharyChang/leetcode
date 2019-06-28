package leetcode

import "strings"

type MapSum_2 struct {
	data map[string]int
}

/** Initialize your data structure here. */
func Constructor_2() MapSum_2 {
	return MapSum_2{
		data: make(map[string]int, 0),
	}
}

func (this *MapSum_2) Insert(key string, val int) {
	this.data[key] = val
}

func (this *MapSum_2) Sum(prefix string) int {
	sum := 0
	for k, v := range this.data {
		if strings.HasPrefix(k, prefix) {
			sum += v
		}
	}
	return sum
}
