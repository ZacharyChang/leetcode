package leetcode

import "fmt"

func getHint(secret string, guess string) string {
	secretMap := str2Map(secret)
	guessMap := str2Map(guess)
	bullCount := countBull(secret, guess)
	dupCount := 0
	for k := range secretMap {
		dupCount += min(secretMap[k], guessMap[k])
	}
	return fmt.Sprintf("%dA%dB", bullCount, dupCount-bullCount)
}

func countBull(secret string, guess string) int {
	res := 0
	for k := range secret {
		if secret[k] == guess[k] {
			res++
		}
	}
	return res
}

func str2Map(str string) map[rune]int {
	res := make(map[rune]int, 0)
	for _, v := range str {
		res[v]++
	}
	return res
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
