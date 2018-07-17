package leetcode

import "testing"

func Test_wiggleMaxLength_dp(t *testing.T) {
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
					1, 17, 5, 10, 13, 15, 10, 5, 16, 8,
				},
			},
			7,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 2, 3, 4, 5, 6, 7, 8, 9,
				},
			},
			2,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					0, 0,
				},
			},
			1,
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					1,
				},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wiggleMaxLength_dp(tt.args.nums); got != tt.want {
				t.Errorf("wiggleMaxLength_dp() = %v, want %v", got, tt.want)
			}
		})
	}
}
