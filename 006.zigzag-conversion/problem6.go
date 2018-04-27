/*

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:

string convert(string text, int nRows);
convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".

*/

package leetcode

import (
	"strings"
)

//-----------------------------------------
// Convert a string to the zigzag pattern.
//-----------------------------------------
func convert(s string, numRows int) string {
	// directly return the input string when numRows is 1
	if numRows == 1 {
		return s
	}
	factor := numRows*2 - 2
	var num int
	result := ""

	array := make([]string, numRows)
	for i, ch := range strings.Split(s, "") {
		num = i % factor
		if num >= numRows {
			num = factor - num
		}
		array[num] += ch

	}
	for _, n := range array {
		result += n
	}
	return result
}
