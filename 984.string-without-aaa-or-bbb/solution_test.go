package leetcode

import (
	"strings"
	"testing"
)

func Test_strWithout3a3b(t *testing.T) {
	type args struct {
		A int
		B int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"[Test Case 1]",
			args{
				1, 2,
			},
		},
		{
			"[Test Case 2]",
			args{
				4, 1,
			},
		},
		{
			"[Test Case 3]",
			args{
				5, 3,
			},
		},
		{
			"[Test Case 4]",
			args{
				5, 11,
			},
		},
		{
			"[Test Case 5]",
			args{
				6, 6,
			},
		},
		{
			"[Test Case 6]",
			args{
				3, 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strWithout3a3b(tt.args.A, tt.args.B); !checkResult(got, tt.args.A, tt.args.B) {
				t.Errorf("strWithout3a3b() = %v is not valid", got)
			}
		})
	}
}

func checkResult(str string, A int, B int) bool {
	if strings.Contains(str, "aaa") {
		return false
	}
	if strings.Contains(str, "bbb") {
		return false
	}
	countMap := make(map[rune]int, 0)
	for _, v := range str {
		countMap[v]++
	}
	return countMap['a'] == A && countMap['b'] == B
}
