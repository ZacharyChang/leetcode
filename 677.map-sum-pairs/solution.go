package leetcode

// Prefix Hashtable
type MapSum struct {
	data   map[string]int
	sumMap map[string]int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		data:   make(map[string]int, 0),
		sumMap: make(map[string]int, 0),
	}
}

func (this *MapSum) Insert(key string, val int) {
	if this.data[key] > 0 {
		for i := 0; i < len(key); i++ {
			this.sumMap[key[:i+1]] = this.sumMap[key[:i+1]] - this.data[key] + val
		}
	} else {
		for i := 0; i < len(key); i++ {
			this.sumMap[key[:i+1]] += val
		}
	}
	this.data[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	return this.sumMap[prefix]
}
