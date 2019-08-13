package leetcode

func reverseOnlyLetters(S string) string {
	bytes := []byte(S)
	start, end := 0, len(S)-1
	for start < end {
		if !isLetter(bytes[start]) {
			start++
			continue
		}
		if !isLetter(bytes[end]) {
			end--
			continue
		}
		bytes[start], bytes[end] = bytes[end], bytes[start]
		start, end = start+1, end-1
	}
	return string(bytes)
}

func isLetter(r byte) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}
