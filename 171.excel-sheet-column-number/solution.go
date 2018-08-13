package leetcode

func titleToNumber(s string) int {
	if len(s) == 1 {
		return int(s[0]) - 'A' + 1
	}
	return titleToNumber(s[:len(s)-1])*26 + titleToNumber(s[len(s)-1:])
}
