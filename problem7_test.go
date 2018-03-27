package leetcode

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				123,
			},
			321,
		},
		{
			"[Test Case 2]",
			args{
				-345,
			},
			-543,
		},
		{
			"[Test Case 3]",
			args{
				1534236469,
			},
			0,
		},
		{
			"[Test Case 4]",
			args{
				1563847412,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.num); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
