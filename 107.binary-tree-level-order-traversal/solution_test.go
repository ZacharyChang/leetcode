package leetcode

import (
	"reflect"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
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
				{15, 7},
				{9, 20},
				{3},
			},
		},
		{
			"[Test Case 2]",
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
		{
			"[Test Case 3]",
			args{
				nil,
			},
			[][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
