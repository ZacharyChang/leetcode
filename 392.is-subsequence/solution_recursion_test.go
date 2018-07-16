package leetcode

import "testing"

func Test_isSubsequence_bf(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence_bf(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence_bf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSubsequence_strings(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence_strings(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence_strings() = %v, want %v", got, tt.want)
			}
		})
	}
}
