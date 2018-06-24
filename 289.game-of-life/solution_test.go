package leetcode

import "testing"

func Test_gameOfLife(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"[Test Case 1]",
			args{
				[][]int{
					{0, 1, 0},
					{0, 0, 1},
					{1, 1, 1},
					{0, 0, 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameOfLife(tt.args.board)
		})
	}
}
