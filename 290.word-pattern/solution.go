package leetcode

import (
	"strings"
)

func wordPattern(pattern string, str string) bool {
	list := strings.Split(str, " ")
	if len(pattern) != len(list) {
		return false
	}
	patternMap := make(map[byte]string, 0)
	strMap := make(map[string]byte, 0)
	for i := 0; i < len(pattern); i++ {
		val, exist := patternMap[pattern[i]]
		if exist {
			if val != list[i] {
				return false
			}
		} else {
			patternMap[pattern[i]] = list[i]
		}
		strVal, exist := strMap[list[i]]
		if exist {
			if strVal != pattern[i] {
				return false
			}
		} else {
			strMap[list[i]] = pattern[i]
		}
	}
	return true
}
