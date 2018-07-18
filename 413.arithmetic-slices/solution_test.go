package leetcode

import "testing"

func Test_numberOfArithmeticSlices(t *testing.T) {
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
			if got := numberOfArithmeticSlices(tt.args.A); got != tt.want {
				t.Errorf("numberOfArithmeticSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
