package leetcode

import "testing"

func Test_isRectangleOverlap(t *testing.T) {
	type args struct {
		rec1 []int
		rec2 []int
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
					0, 0, 2, 2,
				},
				[]int{
					1, 1, 3, 3,
				},
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					0, 0, 1, 1,
				},
				[]int{
					1, 0, 2, 1,
				},
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					1, 1, 2, 2,
				},
				[]int{
					1, 1, 3, 3,
				},
			},
			true,
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					1, 1, 2, 2,
				},
				[]int{
					1, 1, 2, 2,
				},
			},
			true,
		},
		{
			"[Test Case 5]",
			args{
				[]int{
					2, 17, 6, 20,
				},
				[]int{
					3, 8, 6, 20,
				},
			},
			true,
		},
		{
			"[Test Case 6]",
			args{
				[]int{
					-687153884, -854669644, -368882013, -788694078,
				},
				[]int{
					540420176, -909203694, 655002739, -577226067,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRectangleOverlap(tt.args.rec1, tt.args.rec2); got != tt.want {
				t.Errorf("isRectangleOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
