package leetcode

import "testing"

func Test_peakIndexInMountainArray(t *testing.T) {
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
					0, 1, 2, 1,
				},
			},
			2,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					0, 1, 2,
				},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexInMountainArray(tt.args.A); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
