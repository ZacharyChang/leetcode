// Given a digit string, return all possible letter combinations that the number could represent.

// A mapping of digit to letters (just like on the telephone buttons) is given below.

// Input:Digit string "23"
// Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
// Note:
// Although the above answer is in lexicographical order, your answer could be in any order you want.
package leetcode

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	array := make(map[string]string)
	result := []string{""}
	array["2"] = "abc"
	array["3"] = "def"
	array["4"] = "ghi"
	array["5"] = "jkl"
	array["6"] = "mno"
	array["7"] = "pqrs"
	array["8"] = "tuv"
	array["9"] = "wxyz"

	for i := 0; i < len(digits); i++ {
		temp := array[string(digits[i])]
		// make a copy of the last result
		resultCopy := make([]string, len(result))
		copy(resultCopy, result)
		// important, iterator from the result itself
		for k := 0; k < len(resultCopy); k++ {
			for j := 0; j < len(temp); j++ {
				result = append(result, resultCopy[k]+string(temp[j]))
				fmt.Println(resultCopy[k] + string(temp[j]))
			}
		}
		// remove old elements
		result = result[len(resultCopy):]
	}
	return result
}
