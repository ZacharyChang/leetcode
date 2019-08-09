package leetcode

import "strings"

func uniqueMorseRepresentations(words []string) int {
	uniqueMorse := make(map[string]int, 0)
	charArray := []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..",
	}
	for _, word := range words {
		var morse strings.Builder
		for _, ch := range word {
			morse.Write([]byte(charArray[ch-'a']))
		}
		uniqueMorse[morse.String()]++
	}

	return len(uniqueMorse)
}
