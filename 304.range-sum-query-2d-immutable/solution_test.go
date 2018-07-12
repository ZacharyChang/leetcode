package leetcode

import "testing"

func TestNumMatrix_SumRegion(t *testing.T) {
	type args struct {
		row1 int
		col1 int
		row2 int
		col2 int
	}
	tests := []struct {
		name   string
		matrix NumMatrix
		args   args
		want   int
	}{
		{
			"[Test Case 1]",
			Constructor([][]int{
				{3, 0, 1, 4, 2},
				{5, 6, 3, 2, 1},
				{1, 2, 0, 1, 5},
				{4, 1, 0, 1, 7},
				{1, 0, 3, 0, 5},
			}),
			args{
				2, 1, 4, 3,
			},
			8,
		},
		{
			"[Test Case 2]",
			Constructor([][]int{
				{3, 0, 1, 4, 2},
				{5, 6, 3, 2, 1},
				{1, 2, 0, 1, 5},
				{4, 1, 0, 1, 7},
				{1, 0, 3, 0, 5},
			}),
			args{
				1, 1, 2, 2,
			},
			11,
		},
		{
			"[Test Case 3]",
			Constructor([][]int{
				{3, 0, 1, 4, 2},
				{5, 6, 3, 2, 1},
				{1, 2, 0, 1, 5},
				{4, 1, 0, 1, 7},
				{1, 0, 3, 0, 5},
			}),
			args{
				1, 2, 2, 4,
			},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &NumMatrix{
				array: tt.matrix.array,
				sum:   tt.matrix.sum,
			}
			if got := this.SumRegion(tt.args.row1, tt.args.col1, tt.args.row2, tt.args.col2); got != tt.want {
				t.Errorf("NumMatrix.SumRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}
