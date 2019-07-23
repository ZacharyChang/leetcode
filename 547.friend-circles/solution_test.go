package leetcode

import "testing"

func Test_findCircleNum(t *testing.T) {
	type args struct {
		M [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				[][]int{
					{1, 1, 0},
					{1, 1, 0},
					{0, 0, 1},
				},
			},
			2,
		},
		{
			"[Test Case 2]",
			args{
				[][]int{
					{1, 1, 0},
					{1, 1, 1},
					{0, 1, 1},
				},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCircleNum(tt.args.M); got != tt.want {
				t.Errorf("findCircleNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
