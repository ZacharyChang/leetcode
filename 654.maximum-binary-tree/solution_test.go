package leetcode

import (
	"reflect"
	"testing"
)

func Test_constructMaximumBinaryTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"[Test Case 1]",
			args{
				[]int{
					3, 2, 1, 6, 0, 5,
				},
			},
			&TreeNode{
				6,
				&TreeNode{
					3,
					nil,
					&TreeNode{
						2,
						nil,
						&TreeNode{
							1,
							nil,
							nil,
						},
					},
				},
				&TreeNode{
					5,
					&TreeNode{
						0,
						nil,
						nil,
					},
					nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructMaximumBinaryTree(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructMaximumBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
