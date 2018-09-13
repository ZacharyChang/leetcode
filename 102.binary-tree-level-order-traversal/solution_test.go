package leetcode

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"[Test Case 1]",
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
			[][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
		{
			"[Test Case 1]",
			args{
				&TreeNode{
					3,
					nil,
					nil,
				},
			},
			[][]int{
				{3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
