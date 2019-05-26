package leetcode

import (
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"[Test Case 1]",
			args{
				&ListNode{
					0,
					&ListNode{
						1,
						&ListNode{
							2,
							nil,
						},
					},
				},
				4,
			},
			&ListNode{
				2,
				&ListNode{
					0,
					&ListNode{
						1,
						nil,
					},
				},
			},
		},
		{
			"[Test Case 2]",
			args{
				&ListNode{
					1,
					&ListNode{
						2,
						&ListNode{
							3,
							&ListNode{
								4,
								&ListNode{
									5,
									nil,
								},
							},
						},
					},
				},
				2,
			},
			&ListNode{
				4,
				&ListNode{
					5,
					&ListNode{
						1,
						&ListNode{
							2,
							&ListNode{
								3,
								nil,
							},
						},
					},
				},
			},
		},
		{
			"[Test Case 3]",
			args{
				&ListNode{
					0,
					&ListNode{
						1,
						&ListNode{
							2,
							nil,
						},
					},
				},
				1,
			},
			&ListNode{
				2,
				&ListNode{
					0,
					&ListNode{
						1,
						nil,
					},
				},
			},
		},
		{
			"[Test Case 4]",
			args{
				&ListNode{
					0,
					&ListNode{
						1,
						&ListNode{
							2,
							nil,
						},
					},
				},
				2,
			},
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						0,
						nil,
					},
				},
			},
		},
		{
			"[Test Case 5]",
			args{
				&ListNode{
					0,
					&ListNode{
						1,
						&ListNode{
							2,
							nil,
						},
					},
				},
				3,
			},
			&ListNode{
				0,
				&ListNode{
					1,
					&ListNode{
						2,
						nil,
					},
				},
			},
		},
		{
			"[Test Case 6]",
			args{
				&ListNode{
					1,
					nil,
				},
				3,
			},
			&ListNode{
				1,
				nil,
			},
		},
		{
			"[Test Case 7]",
			args{
				&ListNode{
					1,
					&ListNode{
						2,
						nil,
					},
				},
				10000000000,
			},
			&ListNode{
				1,
				&ListNode{
					2,
					nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
