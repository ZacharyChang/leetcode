// Given a digit string, return all possible letter combinations that the number could represent.

// A mapping of digit to letters (just like on the telephone buttons) is given below.

// Input:Digit string "23"
// Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
// Note:
// Although the above answer is in lexicographical order, your answer could be in any order you want.
package leetcode

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
		// make a empty slice and copy new elements to it
		newResult := []string{}
		for k := 0; k < len(result); k++ {
			for j := 0; j < len(temp); j++ {
				newResult = append(newResult, result[k]+string(temp[j]))
			}
		}
		result = newResult
	}
	return result
}
