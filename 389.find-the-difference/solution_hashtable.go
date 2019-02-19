package leetcode

func findTheDifference_2(s string, t string) byte {
	sMap := str2Map(s)
	tMap := str2Map(t)
	for k := range tMap {
		if tMap[k] != sMap[k] {
			return k
		}
	}
	return '0'
}

func str2Map(s string) map[byte]int {
	res := make(map[byte]int, 0)
	for _, v := range s {
		res[byte(v)]++
	}
	return res
}
