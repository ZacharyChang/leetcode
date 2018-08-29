package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func Test_intersection_2(t *testing.T) {
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
				2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersection_2(tt.args.nums1, tt.args.nums2)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
