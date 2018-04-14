// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

// For example, given n = 3, a solution set is:

// [
//   "((()))",
//   "(()())",
//   "(())()",
//   "()(())",
//   "()()()"
// ]
package leetcode

func generateParenthesis(n int) []string {
	left := 0
	right := 0
	var result []string
	for left < right {
		if left == n && right == n {

		}
	}
	return result
}

/*

func getParenthesis(result *[]string, leftCount int, rightCount int, str string, n int) {
	if leftCount == n && rightCount == n {
		result = &append(result, str)
		return
	}
	if leftCount < n {
		getParenthesis(result, leftCount+1, rightCount, "(")
	}
}
*/
