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

func (d DSU) union(x int, y int) {
	rootX := d.findRoot(x)
	rootY := d.findRoot(y)
	if rootX == rootY {
		return
	}
	d.parent[rootX] = rootY
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	d := NewDSU(len(grid) * len(grid[0]))
	col := len(grid[0])
	for i := 0; i < len(grid); i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				// check if there is adjacent land on the top or left
				if i-1 >= 0 && grid[i-1][j] == '1' {
					d.union(getCount(i, j, col), getCount(i-1, j, col))
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					d.union(getCount(i, j, col), getCount(i, j-1, col))
				}
			} else if grid[i][j] == '0' {
				// set the water grid to -1
				// exclude these grids when count the land grid
				d.parent[getCount(i, j, col)] = -1
			}
		}
	}
	group := make(map[int]int, 0)
	for i := 0; i < len(d.parent); i++ {
		if d.parent[i] >= 0 {
			group[d.findRoot(i)]++
		}
	}
	return len(group)
}

func getCount(i int, j int, col int) int {
	return i*col + j
}
