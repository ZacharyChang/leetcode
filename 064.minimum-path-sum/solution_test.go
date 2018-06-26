package leetcode

import "testing"

func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
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
					{
						1, 3, 1,
					},
					{
						1, 5, 1,
					},
					{
						4, 2, 1,
					},
				},
			},
			7,
		},
		{
			"[Test Case 2]",
			args{
				[][]int{},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
