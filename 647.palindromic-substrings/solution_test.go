package leetcode

import "testing"

func Test_countSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				"abc",
			},
			3,
		},
		{
			"[Test Case 2]",
			args{
				"aba",
			},
			4,
		},
		{
			"[Test Case 3]",
			args{
				"aaa",
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubstrings(tt.args.s); got != tt.want {
				t.Errorf("countSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
