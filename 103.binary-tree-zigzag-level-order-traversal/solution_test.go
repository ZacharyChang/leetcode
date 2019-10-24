package leetcode

import (
	"reflect"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"[Test Case 1]",
			args{
				&TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						20,
						&TreeNode{
							Val: 15,
						},
						&TreeNode{
							Val: 7,
						},
					},
				},
			},
			[][]int{
				{3},
				{20, 9},
				{15, 7},
			},
		},
		{
			"[Test Case 2]",
			args{
				nil,
			},
			[][]int{},
		},
		{
			"[Test Case 3]",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val:  3,
						Left: nil,
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
			},
			[][]int{
				{1},
				{3, 2},
				{4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
