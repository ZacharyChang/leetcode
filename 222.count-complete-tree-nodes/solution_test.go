package leetcode

import "testing"

func Test_countNodes(t *testing.T) {
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
						&TreeNode{
							6,
							nil,
							nil,
						},
						nil,
					},
				},
			},
			6,
		},
		{
			"[Test Case 2]",
			args{
				&TreeNode{
					1,
					&TreeNode{
						2,
						&TreeNode{
							4,
							&TreeNode{
								8,
								nil,
								nil,
							},
							&TreeNode{
								9,
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
						3,
						&TreeNode{
							6,
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
			9,
		},
		{
			"[Test Case 3]",
			args{
				nil,
			},
			0,
		},
		{
			"[Test Case 4]",
			args{
				&TreeNode{
					1,
					nil,
					nil,
				},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodes(tt.args.root); got != tt.want {
				t.Errorf("countNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
