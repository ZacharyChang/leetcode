package leetcode

import "strings"

func canConstruct_loop(ransomNote string, magazine string) bool {
	tmp := magazine
	for i := 0; i < len(ransomNote); i++ {
		index := strings.IndexByte(tmp, ransomNote[i])
		if index == -1 {
			return false
		} else {
			tmp = tmp[:index] + tmp[index+1:]
		}
	}
	return true
}
