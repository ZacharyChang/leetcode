package leetcode

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				[][]int{
					{
						1, 3, 5, 7,
					},
					{
						10, 11, 16, 20,
					},
					{
						23, 30, 34, 50,
					},
				},
				3,
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				[][]int{},
				1,
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				[][]int{
					{1}, {3}, {5},
				},
				0,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
