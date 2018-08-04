package leetcode

func judgeCircle(moves string) bool {
	horizontal, vertical := 0, 0
	for _, v := range moves {
		if v == 'U' {
			vertical++
		} else if v == 'D' {
			vertical--
		} else if v == 'L' {
			horizontal--
		} else if v == 'R' {
			horizontal++
		}
	}
	if horizontal == 0 && vertical == 0 {
		return true
	}
	return false
}
