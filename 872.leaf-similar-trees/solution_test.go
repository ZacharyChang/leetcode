package leetcode

import "testing"

func Test_leafSimilar(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
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
					3,
					&TreeNode{
						5,
						nil,
						nil,
					},
					&TreeNode{
						1,
						nil,
						nil,
					},
				},
				&TreeNode{
					3,
					&TreeNode{
						2,
						&TreeNode{
							5,
							nil,
							nil,
						},
						nil,
					},
					&TreeNode{
						1,
						nil,
						nil,
					},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leafSimilar(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
