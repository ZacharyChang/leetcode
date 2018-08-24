package leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
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
					2, 7, 11, 15,
				},
				9,
			},
			[]int{
				1, 2,
			},
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					2, 7, 11, 15,
				},
				10,
			},
			[]int{
				0, 0,
			},
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					2, 7, 11, 15,
				},
				1,
			},
			[]int{
				0, 0,
			},
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					1, 3, 4,
				},
				6,
			},
			[]int{
				0, 0,
			},
		},
		{
			"[Test Case 5]",
			args{
				[]int{
					2, 3, 7, 9, 11, 15,
				},
				13,
			},
			[]int{
				1, 5,
			},
		},
		{
			"[Test Case 6]",
			args{
				[]int{
					2, 3, 7, 9, 11, 15,
				},
				17,
			},
			[]int{
				1, 6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
