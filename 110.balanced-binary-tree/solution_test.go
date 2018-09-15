package leetcode

import (
	"reflect"
	"testing"
)

func Test_isBalanced(t *testing.T) {
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
				nil,
			},
			true,
		},
		{
			"[Test Case 2]",
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
			true,
		},
		{
			"[Test Case 3]",
			args{
				&TreeNode{
					1,
					&TreeNode{
						2,
						&TreeNode{
							3,
							&TreeNode{
								4,
								nil,
								nil,
							},
							nil,
						},
						nil,
					},
					&TreeNode{
						2,
						nil,
						&TreeNode{
							3,
							nil,
							&TreeNode{
								4,
								nil,
								nil,
							},
						},
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
