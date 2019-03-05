package leetcode

import "testing"

func Test_subarraySum1(t *testing.T) {
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
					1, 1, 1,
				},
				2,
			},
			2,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 2, 3,
				},
				3,
			},
			2,
		},
		{
			"[Test Case 3]",
			args{
				[]int{},
				2,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum1(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum_1() = %v, want %v", got, tt.want)
			}
		})
	}
}
