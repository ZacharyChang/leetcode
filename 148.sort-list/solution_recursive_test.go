package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortList_2(t *testing.T) {
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
					4,
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
			},
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						&ListNode{
							4,
							nil,
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
						3,
						&ListNode{
							4,
							&ListNode{
								2,
								nil,
							},
						},
					},
				},
			},
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						&ListNode{
							4,
							nil,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList_2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
