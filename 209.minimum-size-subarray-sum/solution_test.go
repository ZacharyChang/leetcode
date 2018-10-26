package leetcode

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		s    int
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
				7,
				[]int{
					2, 3, 1, 2, 4, 3,
				},
			},
			2,
		},
		{
			"[Test Case 2]",
			args{
				3,
				[]int{
					1, 1,
				},
			},
			0,
		},
		{
			"[Test Case 1]",
			args{
				7,
				[]int{
					2, 3, 1, 2, 4, 3, 2, 1,
				},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.s, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
