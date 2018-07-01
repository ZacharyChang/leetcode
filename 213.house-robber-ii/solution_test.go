package leetcode

import "testing"

func Test_rob(t *testing.T) {
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
				[]int{},
			},
			0,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					2, 3, 2,
				},
			},
			3,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					1, 2, 3, 1,
				},
			},
			4,
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
		{
			"[Test Case 5]",
			args{
				[]int{
					1, 2,
				},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
