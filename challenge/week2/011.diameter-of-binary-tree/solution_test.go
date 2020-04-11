package leetcode

import "testing"

func Test_diameterOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			3,
		},
		{
			"[Test Case 2]",
			args{
				nil,
			},
			0,
		},
		{
			"[Test Case 3]",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
			1,
		},
		{
			"[Test Case 4]",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 6,
							},
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
			},
			4,
		},
		{
			"[Test Case 5]",
			args{
				&TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: -7,
					},
					Right: &TreeNode{
						Val: -3,
						Left: &TreeNode{
							Val: -9,
							Left: &TreeNode{
								Val: 9,
								Left: &TreeNode{
									Val: 6,
									Left: &TreeNode{
										Val: 0,
										Right: &TreeNode{
											Val: -1,
										},
									},
									Right: &TreeNode{
										Val: 6,
										Left: &TreeNode{
											Val: -4,
										},
									},
								},
							},
							Right: &TreeNode{
								Val: -7,
								Left: &TreeNode{
									Val: -6,
									Left: &TreeNode{
										Val: 5,
									},
								},
								Right: &TreeNode{
									Val: -6,
									Left: &TreeNode{
										Val: 9,
										Left: &TreeNode{
											Val: -2,
										},
									},
								},
							},
						},
						Right: &TreeNode{
							Val: -3,
							Left: &TreeNode{
								Val: -4,
							},
						},
					},
				},
			},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
