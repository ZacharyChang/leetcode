package leetcode

import "testing"

func Test_minCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
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
					10, 15, 20,
				},
			},
			15,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 100, 1, 1, 1, 100, 1, 1, 100, 1,
				},
			},
			6,
		},
		{
			"[Test Case 3]",
			args{
				[]int{},
			},
			0,
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					1,
				},
			},
			1,
		},
		{
			"[Test Case 5]",
			args{
				[]int{
					2, 1,
				},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
