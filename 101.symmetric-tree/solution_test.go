package leetcode

import (
	"reflect"
	"testing"
)

func Test_isSymmetric(t *testing.T) {
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
					1,
					&TreeNode{
						2,
						&TreeNode{
							3,
							nil,
							nil,
						},
						&TreeNode{
							4,
							nil,
							nil,
						},
					},
					&TreeNode{
						2,
						&TreeNode{
							4,
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
			},
			true,
		},
		{
			"[Test Case 2]",
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
			false,
		},
		{
			"[Test Case 3]",
			args{
				&TreeNode{
					1,
					&TreeNode{
						2,
						nil,
						nil,
					},
					nil,
				},
			},
			false,
		},
		{
			"[Test Case 4]",
			args{
				&TreeNode{
					1,
					&TreeNode{
						2,
						&TreeNode{
							3,
							nil,
							nil,
						},
						&TreeNode{
							4,
							nil,
							nil,
						},
					},
					&TreeNode{
						2,
						&TreeNode{
							4,
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
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
