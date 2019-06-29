package leetcode

import "testing"

func Test_checkPossibility(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				[]int{
					4, 2, 3,
				},
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					4, 2, 1,
				},
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					3, 4, 2, 3,
				},
			},
			false,
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					2, 3, 3, 2, 4,
				},
			},
			true,
		},
		{
			"[Test Case 5]",
			args{
				[]int{
					1, 2, 3, 3, 4,
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPossibility(tt.args.nums); got != tt.want {
				t.Errorf("checkPossibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
