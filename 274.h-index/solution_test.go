package leetcode

import "testing"

func Test_hIndex(t *testing.T) {
	type args struct {
		citations []int
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
					3, 0, 6, 1, 5,
				},
			},
			3,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 3, 5, 0, 6,
				},
			},
			3,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					0, 1,
				},
			},
			1,
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					1, 2,
				},
			},
			1,
		},
		{
			"[Test Case 5]",
			args{
				[]int{
					0, 1, 0,
				},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex(tt.args.citations); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
