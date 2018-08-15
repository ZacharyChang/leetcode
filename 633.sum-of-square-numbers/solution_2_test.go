package leetcode

import "testing"

func Test_judgeSquareSum_2(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				5,
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				3,
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				1,
			},
			true,
		},
		{
			"[Test Case 4]",
			args{
				2147483646,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeSquareSum_2(tt.args.c); got != tt.want {
				t.Errorf("judgeSquareSum_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
