package leetcode

import "testing"

func Test_isValid(t *testing.T) {
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
				"({})",
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				"",
			},
			true,
		},
		{
			"[Test Case 3]",
			args{
				"{{()()}[]}",
			},
			true,
		},
		{
			"[Test Case 4]",
			args{
				"{{()()}[}",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
