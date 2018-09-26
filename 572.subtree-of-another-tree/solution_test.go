package leetcode

import "testing"

func Test_isSubtree(t *testing.T) {
	type args struct {
		s *TreeNode
		t *TreeNode
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
						4,
						&TreeNode{
							1,
							nil,
							nil,
						},
						&TreeNode{
							2,
							nil,
							nil,
						},
					},
					&TreeNode{
						5,
						nil,
						nil,
					},
				},
				&TreeNode{
					4,
					&TreeNode{
						1,
						nil,
						nil,
					},
					&TreeNode{
						2,
						nil,
						nil,
					},
				},
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				&TreeNode{
					3,
					&TreeNode{
						4,
						&TreeNode{
							1,
							nil,
							nil,
						},
						&TreeNode{
							2,
							&TreeNode{
								0,
								nil,
								nil,
							},
							nil,
						},
					},
					&TreeNode{
						5,
						nil,
						nil,
					},
				},
				&TreeNode{
					4,
					&TreeNode{
						1,
						nil,
						nil,
					},
					&TreeNode{
						2,
						nil,
						nil,
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
