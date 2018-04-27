package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
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
					nil,
				},
				1,
			},
			nil,
		},
		{
			"[Test Case 2]",
			args{
				&ListNode{
					1,
					&ListNode{
						2,
						nil,
					},
				},
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
					1,
					&ListNode{
						2,
						&ListNode{
							3,
							nil,
						},
					},
				},
				3,
			},
			&ListNode{
				2,
				&ListNode{
					3,
					nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
