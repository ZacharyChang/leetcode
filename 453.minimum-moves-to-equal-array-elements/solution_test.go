package leetcode

import "testing"

func Test_minMoves(t *testing.T) {
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
				[]int{
					1, 2, 3,
				},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoves(tt.args.nums); got != tt.want {
				t.Errorf("minMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
