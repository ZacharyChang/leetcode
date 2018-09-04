package leetcode

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"[Test Case 1]",
			args{
				[]int{
					3, 9, 20, 15, 7,
				},
				[]int{
					9, 3, 15, 20, 7,
				},
			},
			&TreeNode{
				3,
				&TreeNode{
					9,
					nil,
					nil,
				},
				&TreeNode{
					20,
					&TreeNode{
						15,
						nil,
						nil,
					},
					&TreeNode{
						7,
						nil,
						nil,
					},
				},
			},
		},
		{
			"[Test Case 2]",
			args{
				[]int{},
				[]int{},
			},
			nil,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					1, 2,
				},
				[]int{
					2, 1,
				},
			},
			&TreeNode{
				1,
				&TreeNode{
					2,
					nil,
					nil,
				},
				nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
