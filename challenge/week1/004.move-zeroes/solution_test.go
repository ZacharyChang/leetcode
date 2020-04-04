package leetcode

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
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
					0, 1, 0, 3, 12,
				},
			},
			[]int{
				1, 3, 12, 0, 0,
			},
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					0, 0, 1, 2, 0, 7, 0, 8, 0, 0, 0, 10, 0,
				},
			},
			[]int{
				1, 2, 7, 8, 10, 0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					0, 0, 1,
				},
			},
			[]int{
				1, 0, 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveZeroes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}