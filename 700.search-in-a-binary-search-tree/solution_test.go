package leetcode

import (
	"reflect"
	"testing"
)

func Test_searchBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
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
					4,
					&TreeNode{
						2,
						&TreeNode{
							1,
							nil,
							nil,
						},
						&TreeNode{
							3,
							nil,
							nil,
						},
					},
					&TreeNode{
						7,
						nil,
						nil,
					},
				},
				7,
			},
			&TreeNode{
				7,
				nil,
				nil,
			},
		},
		{
			"[Test Case 2]",
			args{
				&TreeNode{
					4,
					&TreeNode{
						2,
						&TreeNode{
							1,
							nil,
							nil,
						},
						&TreeNode{
							3,
							nil,
							nil,
						},
					},
					&TreeNode{
						7,
						nil,
						nil,
					},
				},
				2,
			},
			&TreeNode{
				2,
				&TreeNode{
					1,
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
		{
			"[Test Case 3]",
			args{
				nil,
				2,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
