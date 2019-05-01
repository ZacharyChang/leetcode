package leetcode

import "testing"

func Test_maxArea_2(t *testing.T) {
	type args struct {
		height []int
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
					1, 8, 6, 2, 5, 4, 8, 3, 7,
				},
			},
			49,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 1,
				},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea_2(tt.args.height); got != tt.want {
				t.Errorf("maxArea_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
