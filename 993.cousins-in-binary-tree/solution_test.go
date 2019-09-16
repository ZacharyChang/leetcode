package leetcode

import "testing"

func Test_isCousins(t *testing.T) {
	type args struct {
		root *TreeNode
		x    int
		y    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
				4,
				3,
			},
			false,
		},
		{
			"[Test Case 2]",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  2,
						Left: nil,
						Right: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:  3,
						Left: nil,
						Right: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
					},
				},
				4,
				3,
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				&TreeNode{
					Val:  1,
					Left: nil,
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val:  4,
								Left: nil,
								Right: &TreeNode{
									Val:   6,
									Left:  nil,
									Right: nil,
								},
							},
							Right: nil,
						},
						Right: &TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
					},
				},
				5,
				3,
			},
			false,
		},
		{
			"[Test Case 4]",
			args{
				&TreeNode{
					Val:  1,
					Left: nil,
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val:  4,
								Left: nil,
								Right: &TreeNode{
									Val:   6,
									Left:  nil,
									Right: nil,
								},
							},
							Right: nil,
						},
						Right: &TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
					},
				},
				3,
				5,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCousins(tt.args.root, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("isCousins() = %v, want %v", got, tt.want)
			}
		})
	}
}
