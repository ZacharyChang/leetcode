package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func Test_intersection_3(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
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
					4, 9, 5,
				},
				[]int{
					9, 4, 9, 8, 4,
				},
			},
			[]int{
				4, 9,
			},
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 2, 2, 1,
				},
				[]int{
					2, 2,
				},
			},
			[]int{
				2, 2,
			},
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					1, 2, 2, 1,
				},
				[]int{},
			},
			[]int{},
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					1, 2, 3, 2, 1,
				},
				[]int{2, 2, 3, 4},
			},
			[]int{2, 2, 3},
		},

		{
			"[Test Case 5]",
			args{
				[]int{
					4, 9, 5,
				},
				[]int{
					9, 4, 9, 8, 4,
				},
			},
			[]int{
				4, 9,
			},
		},
		{
			"[Test Case 6]",
			args{
				[]int{
					1, 2, 3, 2, 1, 6,
				},
				[]int{2, 2, 3, 4, 0},
			},
			[]int{2, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersect_2(tt.args.nums1, tt.args.nums2)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
