package leetcode

import "testing"

func Test_tree2str(t *testing.T) {
	type args struct {
		t *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
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
						nil,
					},
					&TreeNode{
						3,
						nil,
						nil,
					},
				},
			},
			"1(2(4))(3)",
		},
		{
			"[Test Case 2]",
			args{
				nil,
			},
			"",
		},
		{
			"[Test Case 3]",
			args{
				&TreeNode{
					1,
					&TreeNode{
						2,
						nil,
						&TreeNode{
							4,
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
			"1(2()(4))(3)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree2str(tt.args.t); got != tt.want {
				t.Errorf("tree2str() = %v, want %v", got, tt.want)
			}
		})
	}
}
