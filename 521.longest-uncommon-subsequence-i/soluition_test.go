package leetcode

import "testing"

func Test_findLUSlength(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				"aba",
				"cdc",
			},
			3,
		},
		{
			"[Test Case 2]",
			args{
				"aba",
				"abc",
			},
			3,
		},
		{
			"[Test Case 3]",
			args{
				"aba",
				"ababbc",
			},
			6,
		},
		{
			"[Test Case 4]",
			args{
				"ab",
				"ab",
			},
			-1,
		},
		{
			"[Test Case 5]",
			args{
				"abb",
				"aa",
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLUSlength(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("findLUSlength() = %v, want %v", got, tt.want)
			}
		})
	}
}
