package leetcode

import "testing"

func Test_numberOfArithmeticSlices_bf(t *testing.T) {
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
				[]int{1, 2, 3, 4},
			},
			3,
		},
		{
			"[Test Case 2]",
			args{
				[]int{1, 2, 3, 8, 9, 10},
			},
			2,
		},
		{
			"[Test Case 3]",
			args{
				[]int{1, 2},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArithmeticSlices_bf(tt.args.A); got != tt.want {
				t.Errorf("numberOfArithmeticSlices_bf() = %v, want %v", got, tt.want)
			}
		})
	}
}
