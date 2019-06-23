package leetcode

import (
	"reflect"
	"testing"
)

func Test_matrixReshape(t *testing.T) {
	type args struct {
		nums [][]int
		r    int
		c    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"[Test Case 1]",
			args{
				[][]int{
					{1, 2},
					{3, 4},
				},
				1,
				4,
			},
			[][]int{
				{
					1, 2, 3, 4,
				},
			},
		},
		{
			"[Test Case 2]",
			args{
				[][]int{
				},
				1,
				4,
			},
			[][]int{
			},
		},
		{
			"[Test Case 3]",
			args{
				[][]int{
					{1, 2},
					{3, 4},
				},
				1,
				3,
			},
			[][]int{
				{1, 2},
				{3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixReshape(tt.args.nums, tt.args.r, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixReshape() = %v, want %v", got, tt.want)
			}
		})
	}
}
