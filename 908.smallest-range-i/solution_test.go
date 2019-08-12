package leetcode

import "testing"

func Test_smallestRangeI(t *testing.T) {
	type args struct {
		A []int
		K int
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
					0, 10,
				},
				2,
			},
			6,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 3, 6,
				},
				3,
			},
			0,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					1,
				},
				0,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestRangeI(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("smallestRangeI() = %v, want %v", got, tt.want)
			}
		})
	}
}
