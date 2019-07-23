package leetcode

import (
	"sort"
)

// DSU is a type of Disjoint Sets UnionFind
// performance improved by using array rather than hashtable
type DSU struct {
	parent []int
}

func NewDSD() DSU {
	// max email number is 1000 * 10
	parent := make([]int, 10000)
	for i := range parent {
		parent[i] = i
	}
	return DSU{
		parent: parent,
	}
}

func (d *DSU) findRoot(x int) int {
	// can use recursive here
	// but use loop because of performance
	for d.parent[x] != x {
		x = d.parent[x]
	}
	return x
}

func (d *DSU) union(a int, b int) bool {
	rootA := d.findRoot(a)
	rootB := d.findRoot(b)
	if rootA == rootB {
		return false
	}
	d.parent[rootA] = rootB
	return true
}

// union find
func accountsMerge(accounts [][]string) [][]string {
	res := make([][]string, 0)
	emailToAccount := make(map[string]string, 0)
	emailToID := make(map[string]int, 0)

	d := NewDSD()

	id := 0
	for i := 0; i < len(accounts); i++ {
		for j := 1; j < len(accounts[i]); j++ {
			if _, ok := emailToID[accounts[i][j]]; !ok {
				emailToID[accounts[i][j]] = id
				id++
			}
			d.union(emailToID[accounts[i][1]], emailToID[accounts[i][j]])
			emailToAccount[accounts[i][j]] = accounts[i][0]
		}
	}

	group := make(map[int][]string, 0)

	// create the group by different root
	for k := range emailToAccount {
		rootID := d.findRoot(emailToID[k])
		group[rootID] = append(group[rootID], k)
	}
	for _, v := range group {
		// in sorted order
		sort.Strings(v)
		res = append(res, append([]string{emailToAccount[v[0]]}, v...))
	}

	return res
}
