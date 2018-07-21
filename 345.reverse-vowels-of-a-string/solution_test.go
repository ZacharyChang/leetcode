package leetcode

import "testing"

func Test_reverseVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"[Test Case 1]",
			args{
				"leetcode",
			},
			"leotcede",
		},
		{
			"[Test Case 2]",
			args{
				"hello",
			},
			"holle",
		},
		{
			"[Test Case 3]",
			args{
				"",
			},
			"",
		},
		{
			"[Test Case 4]",
			args{
				"Aa",
			},
			"aA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.args.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
