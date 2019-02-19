package leetcode

func findTheDifference(s string, t string) byte {
	sum1, sum2 := 0, 0
	for _, v := range s {
		sum1 += int(v)
	}
	for _, v := range t {
		sum2 += int(v)
	}
	return byte(sum2 - sum1)
}
