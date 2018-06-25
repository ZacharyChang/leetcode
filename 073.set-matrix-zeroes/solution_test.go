package leetcode

import "testing"

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"[Test Case 1]",
			args{
				[][]int{
					{
						1, 1, 1,
					},
					{
						1, 0, 1,
					},
					{
						1, 1, 1,
					},
				},
			},
		},
		{
			"[Test Case 2]",
			args{
				[][]int{
					{
						0, 1, 2, 0,
					},
					{
						3, 4, 5, 2,
					},
					{
						1, 3, 1, 5,
					},
				},
			},
		},
		{
			"[Test Case 3]",
			args{
				[][]int{
					{
						1, 1, 1,
					},
					{
						0, 1, 2,
					},
				},
			},
		},
		{
			"[Test Case 4]",
			args{
				[][]int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
		})
	}
}
