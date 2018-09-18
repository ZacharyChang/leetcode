package leetcode

import "testing"

func Test_maxDepth(t *testing.T) {
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
			3,
		},
		{
			"[Test Case 1]",
			args{
				&TreeNode{
					3,
					&TreeNode{
						9,
						nil,
						nil,
					},
					nil,
				},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
