package leetcode

import "testing"

func Test_repeatedSubstringPattern(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				"abcabcabcabc",
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				"aabbaa",
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				"a",
			},
			false,
		},
		{
			"[Test Case 4]",
			args{
				"aabaaba",
			},
			false,
		},
		{
			"[Test Case 5]",
			args{
				"bb",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPattern(tt.args.s); got != tt.want {
				t.Errorf("repeatedSubstringPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
