package leetcode

import "testing"

func Test_findMaxAverage(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"[Test Case 1]",
			args{
				[]int{
					1, 12, -5, -6, 50, 3,
				},
				4,
			},
			12.75,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 2, 3,
				},
				4,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxAverage(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findMaxAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
