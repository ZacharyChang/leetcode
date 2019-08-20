package leetcode

func minAddToMakeValid(S string) int {
	res := 0
	leftParenthesesCount := 0
	for _, ch := range S {
		if ch == '(' {
			leftParenthesesCount++
		} else if ch == ')' {
			if leftParenthesesCount == 0 {
				res++
			} else {
				leftParenthesesCount--
			}
		}
	}
	return res + leftParenthesesCount
}
