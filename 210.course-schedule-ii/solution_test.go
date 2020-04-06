package leetcode

import (
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
		},
		{
			"[Test Case 3]",
			args{
				2,
				[][]int{},
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !checkResult(tt.args.prerequisites, got) {
				t.Errorf("findOrder() = %v, order invalid", got)
			}
		})
	}
}

func checkResult(prerequisites [][]int, order []int) bool {
	orderMap := make(map[int]int, len(order))
	for idx, v := range order {
		orderMap[v] = idx
	}
	for _, pre := range prerequisites {
		if orderMap[pre[0]] < orderMap[pre[1]] {
			return false
		}
	}
	return true
}
