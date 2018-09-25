package leetcode

import "testing"

func Test_findBottomLeftValue_2(t *testing.T) {
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
			if got := findBottomLeftValue_2(tt.args.root); got != tt.want {
				t.Errorf("findBottomLeftValue_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
