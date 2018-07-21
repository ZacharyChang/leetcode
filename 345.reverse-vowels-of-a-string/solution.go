package leetcode

func reverseVowels(s string) string {
	result := make([]byte, len(s))
	copy(result, s)
	m := 0
	n := len(s) - 1
	for m < n {
		if !isVowel(s[m]) {
			m++
		}
		if !isVowel(s[n]) {
			n--
		}
		if isVowel(s[m]) && isVowel(s[n]) {
			result[m], result[n] = result[n], result[m]
			m++
			n--
		}
	}
	return string(result)
}

func isVowel(b byte) bool {
	if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' {
		return true
	}
	if b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U' {
		return true
	}
	return false
}
