package leetcode

import "testing"

func Test_maxRotateFunction_1(t *testing.T) {
	type args struct {
		A []int
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
					4, 3, 2, 6,
				},
			},
			26,
		},
		{
			"[Test Case 2]",
			args{
				[]int{},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRotateFunction_1(tt.args.A); got != tt.want {
				t.Errorf("maxRotateFunction_1() = %v, want %v", got, tt.want)
			}
		})
	}
}
