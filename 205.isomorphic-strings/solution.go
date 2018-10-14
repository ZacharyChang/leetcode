package leetcode

func isIsomorphic(s string, t string) bool {
	stMap := make(map[rune]rune, len(s))
	tsMap := make(map[rune]rune, len(s))

	for i := 0; i < len(s); i++ {
		sVal, tVal := rune(s[i]), rune(t[i])
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
