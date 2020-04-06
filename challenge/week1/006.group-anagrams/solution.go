package leetcode

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	anagramMap := make(map[int][]string)
	for _, str := range strs {
		sum := getSumPoint(str)
		if anagramMap[sum] == nil {
			anagramMap[sum] = []string{str}
		} else {
			anagramMap[sum] = append(anagramMap[sum], str)
		}
	}
	for _, v := range anagramMap {
		res = append(res, v)
	}
	return res
}

func getSumPoint(str string) int {
	res := 1
	for _, b := range str {
		res *= getPoint(byte(b))
	}
	return res
}

var number map[byte]int

func getPoint(b byte) int {
	if len(number) == 0 {
		number = make(map[byte]int)
		c := 0
		for i := 2; number['z'] == 0; i++ {
			isOdd := true
			for j := 2; j*j <= i; j++ {
				if i%j == 0 {
					isOdd = false
					break
				}
			}
			if isOdd {
				number[byte('a'+c)] = i
				c++
			}

		}
	}
	return number[b]
}
