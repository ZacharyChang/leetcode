package leetcode

import "testing"

func Test_isPerfectSquare(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				16,
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				0,
			},
			true,
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
				3,
			},
			false,
		},
		{
			"[Test Case 5]",
			args{
				26,
			},
			false,
		},
		{
			"[Test Case 6]",
			args{
				8,
			},
			false,
		},
		{
			"[Test Case 7]",
			args{
				9,
			},
			true,
		},
		{
			"[Test Case 8]",
			args{
				4,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
