package leetcode

import "testing"

func Test_findMin(t *testing.T) {
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
					4, 5, 6, 7, 0, 1, 2,
				},
			},
			0,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					3, 4, 5, 1, 2,
				},
			},
			1,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					1, 2, 3,
				},
			},
			1,
		},
		{
			"[Test Case 4]",
			args{
				[]int{},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
