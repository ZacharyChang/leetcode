package leetcode

func licenseKeyFormatting(S string, K int) string {
	res := ""
	total := 0
	for i := 0; i < len(S); i++ {
		if S[i] != '-' {
			total++
		}
	}
	cur := total % K
	if cur == 0 {
		cur = K
	}
	for i := 0; i < len(S); i++ {
		if S[i] != '-' {
			if cur == 0 {
				res += "-"
				cur = K
			}
			res += string(toUpper(S[i]))
			cur--
		}
	}
	return res
}

func toUpper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return 'A' + c - 'a'
	}
	return c
}
