package leetcode

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				[]int{
					7, 1, 5, 3, 6, 4,
				},
			},
			7,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 2, 3, 4, 5,
				},
			},
			4,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					7, 6, 4, 3, 1,
				},
			},
			0,
		},
		{
			"[Test Case 4]",
			args{
				[]int{},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
