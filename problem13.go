// Given a roman numeral, convert it to an integer.

// Input is guaranteed to be within the range from 1 to 3999.

package leetcode

func romanToInt(s string) int {
	result := 0
	temp := 0
	romanMap := make(map[byte]int, 0)
	romanMap['I'] = 1
	romanMap['V'] = 5
	romanMap['X'] = 10
	romanMap['L'] = 50
	romanMap['C'] = 100
	romanMap['D'] = 500
	romanMap['M'] = 1000

	for i := len(s) - 1; i >= 0; i-- {
		val := romanMap[s[i]]
		if val < temp {
			result -= val
		} else {
			result += val
		}
		temp = val
	}
	return result
}
