package leetcode

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
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
					1, 3, 5, 6,
				},
				5,
			},
			2,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 3, 5, 6,
				},
				4,
			},
			2,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					1, 3, 5, 6,
				},
				2,
			},
			1,
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					1, 3, 5, 6,
				},
				0,
			},
			0,
		},
		{
			"[Test Case 5]",
			args{
				[]int{
					1, 3, 5, 6,
				},
				8,
			},
			4,
		},
		{
			"[Test Case 6]",
			args{
				[]int{
					1, 3,
				},
				1,
			},
			0,
		},
		{
			"[Test Case 67]",
			args{
				[]int{},
				1,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
