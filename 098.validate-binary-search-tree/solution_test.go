package leetcode

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				&TreeNode{
					2,
					&TreeNode{
						Val: 1,
					},
					&TreeNode{
						Val: 3,
					},
				},
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				&TreeNode{
					5,
					&TreeNode{
						Val: 1,
					},
					&TreeNode{
						4,
						&TreeNode{
							Val: 3,
						},
						&TreeNode{
							Val: 6,
						},
					},
				},
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				&TreeNode{
					1,
					&TreeNode{
						Val: 1,
					},
					nil,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
