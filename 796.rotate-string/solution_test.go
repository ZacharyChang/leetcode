package leetcode

import "testing"

func Test_rotateString(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				"abcde",
				"cdeab",
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				"abcde",
				"abced",
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				"",
				"",
			},
			true,
		},
		{
			"[Test Case 4]",
			args{
				"ohbrwzxvxe",
				"uornhegseo",
			},
			false,
		},
		{
			"[Test Case 5]",
			args{
				"a",
				"ab",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateString(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("rotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
