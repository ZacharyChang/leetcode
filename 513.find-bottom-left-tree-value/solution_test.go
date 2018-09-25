package leetcode

import "testing"

func Test_findBottomLeftValue(t *testing.T) {
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
					2,
					&TreeNode{
						1,
						nil,
						nil,
					},
					&TreeNode{
						3,
						nil,
						nil,
					},
				},
			},
			1,
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
							nil,
							nil,
						},
						nil,
					},
					&TreeNode{
						3,
						&TreeNode{
							5,
							&TreeNode{
								7,
								nil,
								nil,
							},
							nil,
						},
						&TreeNode{
							6,
							nil,
							nil,
						},
					},
				},
			},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBottomLeftValue(tt.args.root); got != tt.want {
				t.Errorf("findBottomLeftValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
