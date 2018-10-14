package leetcode

func isIsomorphic(s string, t string) bool {
	// map[byte]byte can be replaced by array
	stMap := make([]byte, 256)
	tsMap := make([]byte, 256)

	for i := 0; i < len(s); i++ {
		sVal, tVal := s[i], t[i]
		if stMap[sVal] > 0 && stMap[sVal] != tVal {
			return false
		}
		if tsMap[tVal] > 0 && tsMap[tVal] != sVal {
			return false
		}
		stMap[sVal] = tVal
		tsMap[tVal] = sVal
	}
	return true
}
