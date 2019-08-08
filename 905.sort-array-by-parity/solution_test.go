package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParity(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"[Test Case 1]",
			args{
				[]int{
					3, 1, 2, 4,
				},
			},
			[][]int{
				{
					2, 4, 3, 1,
				},
				{
					4, 2, 3, 1,
				},
				{
					2, 4, 1, 3,
				},
				{
					4, 2, 1, 3,
				},
			},
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 2,
				},
			},
			[][]int{
				{
					2, 1,
				},
			},
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					4, 2, 1, 3,
				},
			},
			[][]int{
				{
					2, 4, 3, 1,
				},
				{
					4, 2, 3, 1,
				},
				{
					2, 4, 1, 3,
				},
				{
					4, 2, 1, 3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParity(tt.args.A); !in(got, tt.want) {
				t.Errorf("sortArrayByParity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func in(target []int, results [][]int) bool {
	for _, res := range results {
		if reflect.DeepEqual(target, res) {
			return true
		}
	}
	return false
}
