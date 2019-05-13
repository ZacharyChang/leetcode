package leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
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
				0, 1,
			},
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 3,
				},
				4,
			},
			[]int{
				0, 1,
			},
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					0, 1, 2,
				},
				3,
			},
			[]int{
				1, 2,
			},
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					3, 3,
				},
				6,
			},
			[]int{
				0, 1,
			},
		},
		{
			"[Test Case 5]",
			args{
				[]int{
					3, 2, 3,
				},
				6,
			},
			[]int{
				0, 2,
			},
		},
		{
			"[Test Case 6]",
			args{
				[]int{
					1,
				},
				1,
			},
			[]int{
				0, 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
