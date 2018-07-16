package leetcode

import "testing"

func Test_countNumbersWithUniqueDigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				1,
			},
			10,
		},
		{
			"[Test Case 2]",
			args{
				2,
			},
			91,
		},
		{
			"[Test Case 2]",
			args{
				2,
			},
			91,
		},
		{
			"[Test Case 3]",
			args{
				3,
			},
			739,
		},
		{
			"[Test Case 4]",
			args{
				0,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNumbersWithUniqueDigits(tt.args.n); got != tt.want {
				t.Errorf("countNumbersWithUniqueDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
