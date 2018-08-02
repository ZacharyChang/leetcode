package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(a string, b string) string {
	aLeft, aRight := parseComplexNumber(a)
	bLeft, bRight := parseComplexNumber(b)
	left := aLeft*bLeft - aRight*bRight
	right := aLeft*bRight + aRight*bLeft
	return fmt.Sprintf("%d+%di", left, right)
}

func parseComplexNumber(str string) (int, int) {
	loc := strings.Index(str, "+")
	a, _ := strconv.Atoi(str[:loc])
	b, _ := strconv.Atoi(str[loc+1 : len(str)-1])
	return a, b
}
