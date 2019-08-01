package leetcode

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	m1, n1 := rec1[0], rec1[1]
	m2, n2 := rec1[2], rec1[3]
	x1, y1 := rec2[0], rec2[1]
	x2, y2 := rec2[2], rec2[3]

	if x1 >= m2 || y1 >= n2 || m1 >= x2 || n1 >= y2 {
		return false
	}
	return true
}
