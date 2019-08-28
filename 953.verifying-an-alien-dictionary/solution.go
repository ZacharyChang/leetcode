package leetcode

func isAlienSorted(words []string, order string) bool {
	orderMap := make(map[rune]int, 0)
	for i, v := range order {
		orderMap[v] = i
	}
	for i := 1; i < len(words); i++ {
		if compare(words[i-1], words[i], orderMap) == 1 {
			return false
		}
	}
	return true
}

func compare(worda string, wordb string, orderMap map[rune]int) int {
	i := 0
	for ; i < len(worda) && i < len(wordb); i++ {
		if worda[i] != wordb[i] {
			if orderMap[rune(worda[i])] < orderMap[rune(wordb[i])] {
				return -1
			} else if orderMap[rune(worda[i])] > orderMap[rune(wordb[i])] {
				return 1
			}
		}
	}
	if i == len(worda) && i == len(wordb) {
		return 0
	}
	if i == len(worda) {
		return -1
	}
	return 1
}
