package leetcode

import (
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head *ListNode
		m    int
		n    int
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
				4,
			},
			&ListNode{
				1,
				&ListNode{
					4,
					&ListNode{
						3,
						&ListNode{
							2,
							&ListNode{
								5,
								nil,
							},
						},
					},
				},
			},
		},
		{
			"[Test Case 2]",
			args{
				&ListNode{
					1,
					nil,
				},
				1,
				1,
			},
			&ListNode{
				1,
				nil,
			},
		},
		{
			"[Test Case 3]",
			args{
				&ListNode{
					3,
					&ListNode{
						5,
						nil,
					},
				},
				1,
				2,
			},
			&ListNode{
				5,
				&ListNode{
					3,
					nil,
				},
			},
		},
		{
			"[Test Case 4]",
			args{
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
				1,
				2,
			},
			&ListNode{
				2,
				&ListNode{
					1,
					&ListNode{
						3,
						nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
