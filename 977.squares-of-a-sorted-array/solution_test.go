package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
