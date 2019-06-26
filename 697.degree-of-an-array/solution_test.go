package leetcode

import "testing"

func Test_findShortestSubArray(t *testing.T) {
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
					1, 2, 2, 3, 1,
				},
			},
			2,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 2, 2, 3, 1, 4, 2,
				},
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findShortestSubArray(tt.args.nums); got != tt.want {
				t.Errorf("findShortestSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
