package leetcode

// tow pointers
func backspaceCompare_2(S string, T string) bool {
	sPos := len(S) - 1
	tPos := len(T) - 1
	sSkip, tSkip := 0, 0
	for sPos >= 0 || tPos >= 0 {
		// find the next position
		for sPos >= 0 {
			if S[sPos] == '#' {
				sSkip++
				sPos--
			} else if sSkip > 0 {
				sSkip--
				sPos--
			} else {
				break
			}
		}
		for tPos >= 0 {
			if T[tPos] == '#' {
				tSkip++
				tPos--
			} else if tSkip > 0 {
				tSkip--
				tPos--
			} else {
				break
			}
		}
		if sPos >= 0 && tPos >= 0 {
			if S[sPos] != T[tPos] {
				return false
			}
		} else if sPos < 0 && tPos < 0 {
			return true
		} else if sPos < 0 || tPos < 0 {
			return false
		}
		sPos--
		tPos--
	}
	return true
}
