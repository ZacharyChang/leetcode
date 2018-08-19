package leetcode

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	distMap := make(map[int]int, 0)
	distMap[distance(p1, p2)]++
	distMap[distance(p1, p3)]++
	distMap[distance(p1, p4)]++
	distMap[distance(p2, p3)]++
	distMap[distance(p2, p4)]++
	distMap[distance(p3, p4)]++
	// 判断是否有点重合(距离为0)
	if distMap[0] > 0 {
		return false
	}
	// 如果是正方形，则所有点的距离只包含邻边与对边两种情况
	if len(distMap) == 2 {
		return true
	}
	return false
}

func distance(p1 []int, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}
