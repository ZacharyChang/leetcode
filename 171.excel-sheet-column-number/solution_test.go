package leetcode

import "testing"

func Test_titleToNumber(t *testing.T) {
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
				"A",
			},
			1,
		},
		{
			"[Test Case 2]",
			args{
				"AB",
			},
			28,
		},
		{
			"[Test Case 3]",
			args{
				"ZY",
			},
			701,
		},
		{
			"[Test Case 4]",
			args{
				"ABC",
			},
			731,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.args.s); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
