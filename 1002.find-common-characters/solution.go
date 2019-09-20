package leetcode

func commonChars(A []string) []string {
	res := make([]string, 0)
	if len(A) == 0 {
		return res
	}
	charMap := str2Map(A[0])
	for i := 1; i < len(A); i++ {
		charMap = merge(charMap, str2Map(A[i]))
	}
	for ch, count := range charMap {
		for i := 0; i < count; i++ {
			res = append(res, string(ch))
		}
	}
	return res
}

func str2Map(str string) map[rune]int {
	res := make(map[rune]int, 0)
	for _, v := range str {
		res[v]++
	}
	return res
}

func merge(a map[rune]int, b map[rune]int) map[rune]int {
	res := make(map[rune]int, 0)
	if len(a) > len(b) {
		a, b = b, a
	}
	for k := range a {
		if b[k] > 0 {
			res[k] = min(a[k], b[k])
		}
	}
	return res
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
