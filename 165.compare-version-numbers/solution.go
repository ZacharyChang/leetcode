package leetcode

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	s1 := strings.Split(version1, ".")
	s2 := strings.Split(version2, ".")

	for i := 0; i < max(len(s1), len(s2)); i++ {
		n1 := getOrZero(s1, i)
		n2 := getOrZero(s2, i)
		c := compare(n1, n2)
		if c != 0 {
			return c
		}
	}
	return 0
}

func compare(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func getOrZero(num []string, index int) int {
	if index >= len(num) {
		return 0
	}
	n, _ := strconv.Atoi(num[index])
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
