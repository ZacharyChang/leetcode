package leetcode

import "testing"

func Test_minMoves2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				[]int{1, 2, 3},
			},
			2,
		},
		{
			"[Test Case 2]",
			args{
				[]int{1, 2, 3, 4, 10},
			},
			11,
		},
		{
			"[Test Case 3]",
			args{
				[]int{1, 2, 3, 4, 7},
			},
			8,
		},
		{
			"[Test Case 4]",
			args{
				[]int{1},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoves2(tt.args.nums); got != tt.want {
				t.Errorf("minMoves2() = %v, want %v", got, tt.want)
			}
		})
	}
}
