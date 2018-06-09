package leetcode

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"[Test Case 1]",
			args{
				[]int{
					1, 2, 5, 6, 0, 0, 0,
				},
				4,
				[]int{
					3, 7,
				},
				2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("befroe: ", tt.args.nums1)
			merge_2(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			fmt.Println("after: ", tt.args.nums1)
		})
	}
}
