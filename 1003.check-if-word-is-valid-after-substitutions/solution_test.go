package leetcode

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				"aabcbc",
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				"abcabcababcc",
			},
			true,
		},
		{
			"[Test Case 3]",
			args{
				"abccba",
			},
			false,
		},
		{
			"[Test Case 4]",
			args{
				"cababc",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.S); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
