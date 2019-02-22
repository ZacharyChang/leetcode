package leetcode

// hash table
func findAnagrams_2(s string, p string) []int {
	sumP := sum(p)
	res := make([]int, 0)
	for i := 0; i <= len(s)-len(p); i++ {
		if sum(s[i:i+len(p)]) == sumP {
			if isAnagram(s[i:i+len(p)], p) {
				res = append(res, i)
			}
		}
	}
	return res
}

func sum(s string) int {
	res := 0
	for _, v := range s {
		res += int(v)
	}
	return res
}

func isAnagram(s string, p string) bool {
	mapS := str2Map(s)
	mapP := str2Map(p)
	for k := range mapS {
		if mapS[k] != mapP[k] {
			return false
		}
	}
	return true
}

func str2Map(s string) map[byte]int {
	res := make(map[byte]int, 0)
	for _, v := range s {
		res[byte(v)]++
	}
	return res
}
