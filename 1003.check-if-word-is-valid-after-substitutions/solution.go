package leetcode

func isValid(S string) bool {
	stack := make([]rune, 0)
	for _, v := range S {
		stack = append(stack, v)
		if len(stack) >= 3 && string(stack[len(stack)-3:]) == "abc" {
			stack = stack[:len(stack)-3]
		}
	}
	return len(stack) == 0
}
