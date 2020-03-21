package leetcode

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				2,
				[][]int{
					{1, 0},
				},
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				2,
				[][]int{
					{1, 0},
					{0, 1},
				},
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				3,
				[][]int{
					{1, 0},
					{2, 0},
					{0, 2},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
