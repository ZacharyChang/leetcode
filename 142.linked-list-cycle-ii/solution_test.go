package leetcode

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
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
					nil,
				},
			},
			nil,
		},
		{
			"[Test Case 2]",
			args{
				func() *ListNode {
					head := &ListNode{
						3,
						&ListNode{
							2,
							&ListNode{
								0,
								&ListNode{
									4,
									nil,
								},
							},
						},
					}
					head.Next.Next.Next.Next = head.Next
					return head
				}(),
			},
			func() *ListNode {
				head := &ListNode{
					2,
					&ListNode{
						0,
						&ListNode{
							4,
							nil,
						},
					},
				}
				head.Next.Next.Next = head
				return head
			}(),
		},
		{
			"[Test Case 3]",
			args{
				&ListNode{
					1,
					&ListNode{
						1,
						nil,
					},
				},
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
