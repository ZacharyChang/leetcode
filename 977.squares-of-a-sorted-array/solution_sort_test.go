package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortedSquares_2(t *testing.T) {
	type args struct {
		A []int
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
					-4, -1, 0, 3, 10,
				},
			},
			[]int{
				0, 1, 9, 16, 100,
			},
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					-7, -3, 2, 3, 11,
				},
			},
			[]int{
				4, 9, 9, 49, 121,
			},
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					0, 1, 3, 5,
				},
			},
			[]int{
				0, 1, 9, 25,
			},
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					-7, -3, 0, 0, 0, 2, 3, 11,
				},
			},
			[]int{
				0, 0, 0, 4, 9, 9, 49, 121,
			},
		},
		{
			"[Test Case 5]",
			args{
				[]int{
					-7, -3, 0,
				},
			},
			[]int{
				0, 9, 49,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares_2(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
