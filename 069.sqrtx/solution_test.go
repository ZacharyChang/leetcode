package leetcode

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				4,
			},
			2,
		},
		{
			"[Test Case 2]",
			args{
				8,
			},
			2,
		},
		{
			"[Test Case 3]",
			args{
				1,
			},
			1,
		},
		{
			"[Test Case 4]",
			args{
				0,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
