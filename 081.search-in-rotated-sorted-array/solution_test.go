package leetcode

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				[]int{4, 5, 1, 2, 3},
				1,
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				[]int{4, 5, 1, 2, 3},
				0,
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				[]int{4, 5, 1, 2, 3},
				4,
			},
			true,
		},
		{
			"[Test Case 4]",
			args{
				[]int{4, 5, 1, 2, 3},
				3,
			},
			true,
		},
		{
			"[Test Case 5]",
			args{
				[]int{4, 5, 6, 7, 0, 1, 2},
				0,
			},
			true,
		},
		{
			"[Test Case 6]",
			args{
				[]int{1, 3, 5},
				0,
			},
			false,
		},
		{
			"[Test Case 7]",
			args{
				[]int{2, 5, 6, 0, 0, 1, 2},
				0,
			},
			true,
		},
		{
			"[Test Case 8]",
			args{
				[]int{1, 3, 1, 1, 1},
				3,
			},
			true,
		},
		{
			"[Test Case 9]",
			args{
				[]int{1, 1, 3},
				0,
			},
			false,
		},
		{
			"[Test Case 10]",
			args{
				[]int{},
				0,
			},
			false,
		},
		{
			"[Test Case 11",
			args{
				[]int{
					1, 2, 3, 4, 5, 6, 7,
				},
				2,
			},
			true,
		},
		{
			"[Test Case 12",
			args{
				[]int{
					5, 6, 7, 1, 2, 3, 4,
				},
				3,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
