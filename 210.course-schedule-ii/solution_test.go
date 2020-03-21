package leetcode

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"[Test Case 1]",
			args{
				2,
				[][]int{
					{
						1, 0,
					},
				},
			},
			[]int{
				0, 1,
			},
		},
		{
			"[Test Case 2]",
			args{
				4,
				[][]int{
					{1, 0},
					{2, 0},
					{3, 1},
					{3, 2},
				},
			},
			[]int{
				0, 1, 2, 3,
			},
		},
		{
			"[Test Case 3]",
			args{
				2,
				[][]int{},
			},
			[]int{
				1, 0,
			},
		},
		{
			"[Test Case 4]",
			args{
				3,
				[][]int{
					{0, 2},
					{1, 2},
					{2, 0},
				},
			},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
