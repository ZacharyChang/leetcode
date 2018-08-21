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
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				5,
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				[][]int{
					{1, 3, 5},
				},
				4,
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				[][]int{},
				4,
			},
			false,
		},
		{
			"[Test Case 4]",
			args{
				[][]int{
					{},
				},
				2,
			},
			false,
		},
		{
			"[Test Case 5]",
			args{
				[][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				3,
			},
			true,
		},
		{
			"[Test Case 6]",
			args{
				[][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				7,
			},
			true,
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
