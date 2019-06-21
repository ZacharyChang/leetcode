package leetcode

import "testing"

func Test_diameterOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				&TreeNode{
					1,
					&TreeNode{
						2,
						&TreeNode{
							4,
							nil,
							nil,
						},
						&TreeNode{
							5,
							nil,
							nil,
						},
					},
					&TreeNode{
						3,
						nil,
						nil,
					},
				},
			},
			3,
		},
		{
			"【Test Case 2]",
			args{
				&TreeNode{
					1,
					nil,
					nil,
				},
			},
			0,
		},
		{
			"【Test Case 3]",
			args{
				nil,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
