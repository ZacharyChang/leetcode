package leetcode

import (
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"[Test Case 1]",
			args{
				&TreeNode{
					1,
					&TreeNode{
						2,
						nil,
						&TreeNode{
							5,
							nil,
							nil,
						},
					},
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
			[]int{
				1, 3, 4,
			},
		},
		{
			"[Test Case 2]",
			args{
				nil,
			},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
