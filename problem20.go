// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.

package leetcode

func isValid(s string) bool {
	temp := ""

	for _, ch := range s {
		length := len(temp)
		// ensure the length > 0
		if length == 0 {
			temp += string(ch)
			continue
		}
		if string(temp[length-1]) == "{" && string(ch) == "}" {
			temp = temp[:length-1]
		} else if string(temp[length-1]) == "[" && string(ch) == "]" {
			temp = temp[:length-1]
		} else if string(temp[length-1]) == "(" && string(ch) == ")" {
			temp = temp[:length-1]
		} else {
			temp += string(ch)
		}
	}
	return len(temp) == 0
}
