package leetcode

import "testing"

func Test_findPeakElement(t *testing.T) {
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
					1, 2, 3, 1,
				},
			},
			2,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1,
				},
			},
			0,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					3, 2, 1,
				},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElement(tt.args.nums); got != tt.want {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
