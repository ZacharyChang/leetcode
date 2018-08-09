package leetcode

func toLowerCase(str string) string {
	result := ""
	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			result += string(str[i] + ('a' - 'A'))
		} else {
			result += string(str[i])
		}
	}
	return result
}
