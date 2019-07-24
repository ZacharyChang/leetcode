package leetcode

type DSU struct {
	equalParent []int
}

func NewDSU() DSU {
	parent := make([]int, 26)
	for i := range parent {
		parent[i] = i
	}
	return DSU{
		equalParent: parent,
	}
}

func (d DSU) findRoot(x int) int {
	for d.equalParent[x] != x {
		x = d.equalParent[x]
	}
	return x
}

func (d DSU) union(x int, y int) bool {
	rootX := d.findRoot(x)
	rootY := d.findRoot(y)
	if rootX == rootY {
		return false
	}
	d.equalParent[rootX] = rootY
	return true
}

func (d DSU) checkEqual(x int, y int) bool {
	if d.findRoot(x) == d.findRoot(y) {
		return true
	}
	return false
}

func equationsPossible(equations []string) bool {
	d := NewDSU()
	for _, equation := range equations {
		x, y := equation[0]-'a', equation[3]-'a'
		if equation[1:3] == "==" {
			d.union(int(x), int(y))
		}
	}
	for _, equation := range equations {
		if equation[1:3] == "!=" {
			x, y := equation[0]-'a', equation[3]-'a'
			if d.checkEqual(int(x), int(y)) {
				return false
			}
		}
	}
	return true
}
