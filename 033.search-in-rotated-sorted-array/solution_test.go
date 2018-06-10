package leetcode

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				[]int{4, 5, 1, 2, 3},
				1,
			},
			2,
		},
		{
			"[Test Case 2]",
			args{
				[]int{4, 5, 1, 2, 3},
				0,
			},
			-1,
		},
		{
			"[Test Case 3]",
			args{
				[]int{4, 5, 1, 2, 3},
				4,
			},
			0,
		},
		{
			"[Test Case 4]",
			args{
				[]int{4, 5, 1, 2, 3},
				3,
			},
			4,
		},
		{
			"[Test Case 5]",
			args{
				[]int{4, 5, 6, 7, 0, 1, 2},
				0,
			},
			4,
		},
		{
			"[Test Case 6]",
			args{
				[]int{1, 3, 5},
				0,
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
