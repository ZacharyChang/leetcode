package leetcode

// two pointers
func isLongPressedName(name string, typed string) bool {
	s1, s2 := 0, 0
	for s1 < len(name) && s2 < len(typed) {
		c1, c2 := 0, 0
		if name[s1] != typed[s2] {
			return false
		}
		for s1+1 < len(name) && name[s1+1] == name[s1] {
			s1++
			c1++
		}
		for s2+1 < len(typed) && typed[s2+1] == typed[s2] {
			s2++
			c2++
		}
		if c1 > c2 {
			return false
		}
		s1++
		s2++
	}
	return s1 == len(name) && s2 == len(typed)
}
