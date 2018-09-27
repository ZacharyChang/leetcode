package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		t1 *TreeNode
		t2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"[Test Case 1]",
			args{
				&TreeNode{
					1,
					&TreeNode{
						3,
						&TreeNode{
							5,
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
				&TreeNode{
					2,
					&TreeNode{
						1,
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
						&TreeNode{
							7,
							nil,
							nil,
						},
					},
				},
			},
			&TreeNode{
				3,
				&TreeNode{
					4,
					&TreeNode{
						5,
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
					5,
					nil,
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
			if got := mergeTrees(tt.args.t1, tt.args.t2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
