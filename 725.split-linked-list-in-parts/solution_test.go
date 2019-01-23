package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_splitListToParts(t *testing.T) {
	type args struct {
		root *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want []*ListNode
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
							nil,
						},
					},
				},
				5,
			},
			[]*ListNode{
				&ListNode{
					1,
					nil,
				},
				&ListNode{
					2,
					nil,
				},
				&ListNode{
					3,
					nil,
				},
				nil,
				nil,
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
									&ListNode{
										6,
										&ListNode{
											7,
											&ListNode{
												8,
												&ListNode{
													9,
													&ListNode{
														10,
														nil,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				3,
			},
			[]*ListNode{
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
				&ListNode{
					5,
					&ListNode{
						6,
						&ListNode{
							7,
							nil,
						},
					},
				},
				&ListNode{
					8,
					&ListNode{
						9,
						&ListNode{
							10,
							nil,
						},
					},
				},
			},
		},
		{
			"[Test Case 3]",
			args{
				nil,
				3,
			},
			[]*ListNode{
				nil,
				nil,
				nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitListToParts(tt.args.root, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				print(got)
				t.Errorf("splitListToParts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func print(list []*ListNode) {
	for _, v := range list {
		func(root *ListNode) {
			fmt.Print("list: ")
			for root != nil {
				fmt.Printf("%d->", root.Val)
				root = root.Next
			}
			fmt.Println()
		}(v)
	}
}
