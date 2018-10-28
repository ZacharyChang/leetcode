package leetcode

func convertToTitle(n int) string {
	res := ""
	for n > 0 {
		n--
		b := byte(n%26 + 'A')
		res = string(b) + res
		n = n / 26
	}
	return res
}
