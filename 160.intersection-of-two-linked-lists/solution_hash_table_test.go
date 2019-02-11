package leetcode

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode_2(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	common := &ListNode{
		8,
		&ListNode{
			4,
			&ListNode{
				5,
				nil,
			},
		},
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
					4,
					&ListNode{
						1,
						common,
					},
				},
				&ListNode{
					5,
					&ListNode{
						0,
						&ListNode{
							1,
							common,
						},
					},
				},
			},
			common,
		},
		{
			"[Test Case 2]",
			args{
				&ListNode{
					4,
					&ListNode{
						1,
						nil,
					},
				},
				&ListNode{
					5,
					&ListNode{
						1,
						nil,
					},
				},
			},
			nil,
		},
		{
			"[Test Case 3]",
			args{
				&ListNode{
					1,
					nil,
				},
				nil,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode_2(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
