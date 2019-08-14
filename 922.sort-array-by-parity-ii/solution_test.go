package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParityII(t *testing.T) {
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
					4, 2, 5, 7,
				},
			},
			[][]int{
				{4, 7, 2, 5},
				{2, 5, 4, 7},
				{2, 7, 4, 5},
				{4, 5, 2, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParityII(tt.args.A); !in(got, tt.want) {
				t.Errorf("sortArrayByParityII() = %v, want %v", got, tt.want)
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
