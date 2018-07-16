package leetcode

import "testing"

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				"abc",
				"ahbgdc",
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				"abc",
				"acbddd",
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				"a",
				"",
			},
			false,
		},
		{
			"[Test Case 4]",
			args{
				"",
				"a",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
