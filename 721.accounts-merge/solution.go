package leetcode

import "sort"

// DSU is a type of Disjoint Sets UnionFind
type DSU struct {
	parent map[string]string
}

func NewDSD() DSU {
	return DSU{
		parent: make(map[string]string, 0),
	}
}

func (d *DSU) findRoot(s string) string {
	for d.parent[s] != "" {
		s = d.parent[s]
	}
	return s
}

func (d *DSU) union(a string, b string) bool {
	rootA := d.findRoot(a)
	rootB := d.findRoot(b)
	if rootA == rootB {
		return false
	}
	d.parent[rootA] = rootB
	return true
}

// union find
// TODO: performance
func accountsMerge(accounts [][]string) [][]string {
	res := make([][]string, 0)
	accountMap := make(map[string]string, 0)

	d := NewDSD()

	for i := 0; i < len(accounts); i++ {
		for j := 1; j < len(accounts[i]); j++ {
			d.union(accounts[i][1], accounts[i][j])
			accountMap[accounts[i][j]] = accounts[i][0]
		}
	}

	group := make(map[string][]string, 0)

	for k := range accountMap {
		group[d.findRoot(k)] = append(group[d.findRoot(k)], k)
	}
	for _, v := range group {
		// in sorted order
		sort.Strings(v)
		res = append(res, append([]string{accountMap[v[0]]}, v...))
	}

	return res
}
