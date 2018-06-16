package leetcode

import "testing"

func Test_thirdMax(t *testing.T) {
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
					3, 2, 1,
				},
			},
			1,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					2, 2, 3, 1,
				},
			},
			1,
		},
		{
			"[Test Case 3]",
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
			if got := thirdMax(tt.args.nums); got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
