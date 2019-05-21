package leetcode

import "testing"

func Test_countSubstrings_2(t *testing.T) {
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
		{
			"[Test Case 4]",
			args{
				"aaaaa",
			},
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubstrings_2(tt.args.s); got != tt.want {
				t.Errorf("countSubstrings_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
