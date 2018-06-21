package leetcode

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"[Test Case 1]",
			args{
				[]int{
					1, 2, 3, 4, 5, 6, 7,
				},
				3,
			},
		},
		{
			"[Test Case 1]",
			args{
				[]int{
					-1,
				},
				2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
		})
	}
}
