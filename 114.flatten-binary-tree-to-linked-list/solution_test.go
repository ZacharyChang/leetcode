package leetcode

import (
	"testing"
)

func Test_flatten(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"[Test Case 1]",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.args.root)
			if !checkFlatten(tt.args.root) {
				t.Errorf("flatten() = %v, want a flatten linked list", tt.args.root)
			}
		})
	}
}

func checkFlatten(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		return false
	}
	return checkFlatten(root.Right)
}
