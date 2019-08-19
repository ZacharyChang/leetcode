package leetcode

import "testing"

func Test_hasGroupsSizeX(t *testing.T) {
	type args struct {
		deck []int
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
					1, 2, 3, 4, 4, 3, 2, 1,
				},
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 1, 1, 2, 2, 2, 3, 3,
				},
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					1,
				},
			},
			false,
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					1, 1,
				},
			},
			true,
		},
		{
			"[Test Case 5]",
			args{
				[]int{},
			},
			false,
		},
		{
			"[Test Case 6]",
			args{
				[]int{
					1, 1, 2, 2, 2, 2,
				},
			},
			true,
		},
		{
			"[Test Case 7]",
			args{
				[]int{
					1, 1, 1, 1, 2, 2, 2, 2, 2, 2,
				},
			},
			true,
		},
		{
			"[Test Case 8]",
			args{
				[]int{
					0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7,
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasGroupsSizeX(tt.args.deck); got != tt.want {
				t.Errorf("hasGroupsSizeX() = %v, want %v", got, tt.want)
			}
		})
	}
}
