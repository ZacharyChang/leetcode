package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"[Test Case 1]",
			args{
				[]*ListNode{
					&ListNode{
						1,
						&ListNode{
							2,
							&ListNode{
								4,
								nil,
							},
						},
					},
					&ListNode{
						1,
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
			&ListNode{
				1,
				&ListNode{
					1,
					&ListNode{
						2,
						&ListNode{
							3,
							&ListNode{
								4,
								&ListNode{
									4,
									nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				printList(got)
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
