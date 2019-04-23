package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
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
						4,
						&ListNode{
							3,
							nil,
						},
					},
				},
				3,
			},
			&ListNode{
				1,
				&ListNode{
					4,
					&ListNode{
						3,
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
						4,
						&ListNode{
							3,
							&ListNode{
								2,
								&ListNode{
									5,
									&ListNode{
										2,
										nil,
									},
								},
							},
						},
					},
				},
				3,
			},
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						2,
						&ListNode{
							4,
							&ListNode{
								3,
								&ListNode{
									5,
									nil,
								},
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
					1,
					nil,
				},
				2,
			},
			&ListNode{
				1,
				nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				print(got)
				print(tt.want)
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func print(head *ListNode) {
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
	fmt.Println()
}
