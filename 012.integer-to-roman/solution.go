package leetcode

var romanArray = [][]string{
	{
		"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX",
	},
	{
		"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC",
	},
	{
		"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM",
	},
	{
		"", "M", "MM", "MMM",
	},
}

// Array
func intToRoman(num int) string {
	res := ""
	res += romanArray[3][num/1000%10]
	res += romanArray[2][num/100%10]
	res += romanArray[1][num/10%10]
	res += romanArray[0][num%10]
	return res
}
