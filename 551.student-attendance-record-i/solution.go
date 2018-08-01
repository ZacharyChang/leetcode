package leetcode

func checkRecord(s string) bool {
	countAbsent := 0
	curLate := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			countAbsent++
			curLate = 0
		} else if s[i] == 'L' {
			curLate++
		} else if s[i] == 'P' {
			curLate = 0
		}
		if countAbsent > 1 || curLate > 2 {
			return false
		}
	}
	return true
}
