package leetcode

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"[Test Case  1]",
			args{
				[]int{
					9, 3, 15, 20, 7,
				},
				[]int{
					9, 15, 7, 20, 3,
				},
			},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.inorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
