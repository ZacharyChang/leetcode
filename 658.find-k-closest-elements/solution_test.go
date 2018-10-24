package leetcode

import (
	"reflect"
	"testing"
)

func Test_findClosestElements(t *testing.T) {
	type args struct {
		arr []int
		k   int
		x   int
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
					1, 2, 3, 4, 5,
				},
				4,
				3,
			},
			[]int{
				1, 2, 3, 4,
			},
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 2, 3, 4, 5,
				},
				4,
				-1,
			},
			[]int{
				1, 2, 3, 4,
			},
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					0, 0, 0, 1, 3, 5, 6, 7, 8, 8,
				},
				2,
				2,
			},
			[]int{
				1, 3,
			},
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					0, 1, 1, 1, 2, 3, 6, 7, 8, 9,
				},
				9,
				4,
			},
			[]int{
				0, 1, 1, 1, 2, 3, 6, 7, 8,
			},
		},
		{
			"[Test Case 5]",
			args{
				[]int{
					0, 0, 1, 2, 3, 3, 4, 7, 7, 8,
				},
				3,
				5,
			},
			[]int{
				3, 3, 4,
			},
		},
		{
			"[Test Case 6]",
			args{
				[]int{
					0, 2, 2, 3, 4, 6, 7, 8, 9, 9,
				},
				4,
				5,
			},
			[]int{
				3, 4, 6, 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findClosestElements(tt.args.arr, tt.args.k, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findClosestElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
