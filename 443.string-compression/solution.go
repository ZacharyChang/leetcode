package leetcode

import (
	"strconv"
)

func compress(chars []byte) int {
	curChar := chars[0]
	curCount := 1
	curPos := 0
	for i := 1; i < len(chars); i++ {
		if chars[i] != curChar {
			if curCount == 1 {
				chars[curPos] = curChar
				curPos++
				curCount = 1
				curChar = chars[i]
			} else if curCount > 1 {
				chars[curPos] = curChar
				copy(chars[curPos+1:], strconv.Itoa(curCount))
				curPos += 1 + len(strconv.Itoa(curCount))
				curCount = 1
				curChar = chars[i]
			}
		} else {
			curCount++
		}
	}
	if curCount == 1 {
		chars[curPos] = curChar
		curPos++
	} else if curCount > 1 {
		chars[curPos] = curChar
		copy(chars[curPos+1:], strconv.Itoa(curCount))
		curPos += 1 + len(strconv.Itoa(curCount))
	}
	return curPos
}
