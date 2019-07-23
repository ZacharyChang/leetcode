package leetcode

type DSU struct {
	parent []int
}

func NewDSU(len int) DSU {
	parent := make([]int, len)
	for i := range parent {
		parent[i] = i
	}
	return DSU{
		parent: parent,
	}
}

func (d DSU) findRoot(x int) int {
	for d.parent[x] != x {
		x = d.parent[x]
	}
	return x
}

func (d DSU) union(x int, y int) bool {
	rootX := d.findRoot(x)
	rootY := d.findRoot(y)
	if rootX == rootY {
		return false
	}
	d.parent[rootX] = rootY
	return true
}

func findCircleNum(M [][]int) int {
	d := NewDSU(len(M))
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[0]); j++ {
			if M[i][j] == 1 {
				d.union(i, j)
			}
		}
	}
	circleCountMap := make(map[int]int, 0)
	for i := 0; i < len(M); i++ {
		circleCountMap[d.findRoot(i)]++
	}
	return len(circleCountMap)
}
