package leetcode

import (
	"reflect"
	"testing"
)

func Test_fairCandySwap(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"[Test Case 1]",
			args{
				[]int{
					1, 1,
				},
				[]int{
					2, 2,
				},
			},
			[]int{
				1, 2,
			},
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 2,
				},
				[]int{
					2, 3,
				},
			},
			[]int{
				1, 2,
			},
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					2, 3,
				},
				[]int{
					1, 2,
				},
			},
			[]int{
				2, 1,
			},
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					2, 3,
				},
				[]int{
					1, 100,
				},
			},
			[]int{
				-1, -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fairCandySwap(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fairCandySwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
