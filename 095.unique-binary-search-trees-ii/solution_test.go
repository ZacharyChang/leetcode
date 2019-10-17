package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_generateTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			"[Test Case 1]",
			args{
				1,
			},
			[]*TreeNode{
				{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
		{
			"[Test Case 2]",
			args{
				2,
			},
			[]*TreeNode{
				{
					Val:  1,
					Left: nil,
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
				{
					Val: 2,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		},
		{
			"[Test Case 3]",
			args{
				0,
			},
			[]*TreeNode{},
		},
		{
			"[Test Case 4]",
			args{
				3,
			},
			[]*TreeNode{
				{
					1,
					nil,
					&TreeNode{
						2,
						nil,
						&TreeNode{
							Val: 3,
						},
					},
				},
				{
					1,
					nil,
					&TreeNode{
						3,
						&TreeNode{
							Val: 2,
						},
						nil,
					},
				},
				{
					2,
					&TreeNode{
						Val: 1,
					},
					&TreeNode{
						Val: 3,
					},
				},
				{
					3,
					&TreeNode{
						1,
						nil,
						&TreeNode{
							Val: 2,
						},
					},
					nil,
				},
				{
					3,
					&TreeNode{
						2,
						&TreeNode{
							Val: 1,
						},
						nil,
					},
					nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateTrees(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				fmt.Println(got)
				t.Errorf("generateTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
