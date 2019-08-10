package leetcode

import "testing"

func Test_validMountainArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				[]int{
					2, 1,
				},
			},
			false,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					3, 5, 5,
				},
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					0, 3, 2, 1,
				},
			},
			true,
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.args.A); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
