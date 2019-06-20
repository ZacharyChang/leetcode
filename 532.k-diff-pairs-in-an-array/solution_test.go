package leetcode

import "testing"

func Test_findPairs(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
					3, 1, 4, 1, 5,
				},
				2,
			},
			2,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 2, 3, 4, 5,
				},
				1,
			},
			4,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					1, 3, 1, 5, 4,
				},
				0,
			},
			1,
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					0, 0, 0,
				},
				0,
			},
			1,
		},
		{
			"[Test Case 5]",
			args{
				[]int{
					1, 2, 3, 4, 5,
				},
				-1,
			},
			0,
		},
		{
			"[Test Case 6]",
			args{
				[]int{
					1, 1, 1, 2, 1,
				},
				1,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPairs(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
