package leetcode

import (
	"reflect"
	"testing"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
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
					5,
					&TreeNode{
						4,
						&TreeNode{
							11,
							&TreeNode{
								7,
								nil,
								nil,
							},
							&TreeNode{
								2,
								nil,
								nil,
							},
						},
						nil,
					},
					&TreeNode{
						8,
						&TreeNode{
							13,
							nil,
							nil,
						},
						&TreeNode{
							4,
							&TreeNode{
								5,
								nil,
								nil,
							},
							&TreeNode{
								1,
								nil,
								nil,
							},
						},
					},
				},
				22,
			},
			[][]int{
				{
					5, 4, 11, 2,
				},
				{
					5, 8, 4, 5,
				},
			},
		},
		{
			"[Test Case 2]",
			args{
				&TreeNode{
					-2,
					&TreeNode{
						-3,
						nil,
						nil,
					},
					nil,
				},
				-5,
			},
			[][]int{
				{
					-2, -3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.sum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
