package leetcode

import (
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"[Test Case 1]",
			args{
				nil,
			},
			nil,
		},
		{
			"[Test Case 2]",
			args{
				&TreeNode{
					4,
					&TreeNode{
						2,
						nil,
						nil,
					},
					&TreeNode{
						7,
						nil,
						&TreeNode{
							1,
							nil,
							nil,
						},
					},
				},
			},
			&TreeNode{
				4,
				&TreeNode{
					7,
					&TreeNode{
						1,
						nil,
						nil,
					},
					nil,
				},
				&TreeNode{
					2,
					nil,
					nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
