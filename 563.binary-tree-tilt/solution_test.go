package leetcode

import "testing"

func Test_findTilt(t *testing.T) {
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
						3,
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
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTilt(tt.args.root); got != tt.want {
				t.Errorf("findTilt() = %v, want %v", got, tt.want)
			}
		})
	}
}
