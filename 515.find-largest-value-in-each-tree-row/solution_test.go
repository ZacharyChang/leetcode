package leetcode

import (
	"reflect"
	"testing"
)

func Test_largestValues(t *testing.T) {
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
						3,
						&TreeNode{
							5,
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
						2,
						nil,
						&TreeNode{
							9,
							nil,
							nil,
						},
					},
				},
			},
			[]int{
				1, 3, 9,
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
			if got := largestValues(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
