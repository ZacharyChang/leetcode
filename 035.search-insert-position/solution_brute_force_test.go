package leetcode

import "testing"

func Test_searchInsert_brute_force(t *testing.T) {
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
				7,
			},
			4,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					1, 3, 5, 0,
				},
				0,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert_brute_force(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert_brute_force() = %v, want %v", got, tt.want)
			}
		})
	}
}
