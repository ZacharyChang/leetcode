package leetcode

import "testing"

func Test_pivotIndex(t *testing.T) {
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
					1, 7, 3, 6, 5, 6,
				},
			},
			3,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 2, 3,
				},
			},
			-1,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					1, 3, 5, 10,
				},
			},
			-1,
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					-1000, 1, 2, 3,
				},
			},
			-1,
		},
		{
			"[Test Case 5]",
			args{
				[]int{
					-1, -1, -1, -1, -1, 0,
				},
			},
			2,
		},
		{
			"[Test Case 6]",
			args{
				[]int{},
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.args.nums); got != tt.want {
				t.Errorf("pivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
