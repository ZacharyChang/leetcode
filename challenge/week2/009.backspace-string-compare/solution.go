package leetcode

func backspaceCompare(S string, T string) bool {
	return toStr(S) == toStr(T)
}

func toStr(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if len(res) >= 1 {
				res = res[:len(res)-1]
			}
		} else {
			res += string(s[i])
		}
	}
	return res
}
