package leetcode

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
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
					2,
					&ListNode{
						4,
						&ListNode{
							3,
							nil,
						},
					},
				},
				&ListNode{
					5,
					&ListNode{
						6,
						&ListNode{
							4,
							nil,
						},
					},
				},
			},
			&ListNode{
				7,
				&ListNode{
					0,
					&ListNode{
						8,
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
						8,
						nil,
					},
				},
				&ListNode{
					0,
					nil,
				},
			},
			&ListNode{
				1,
				&ListNode{
					8,
					nil,
				},
			},
		},
		{
			"[Test Case 3]",
			args{
				&ListNode{
					5,
					nil,
				},
				&ListNode{
					5,
					nil,
				},
			},
			&ListNode{
				0,
				&ListNode{
					1,
					nil,
				},
			},
		},
		{
			"[Test Case 4]",
			args{
				&ListNode{
					9,
					&ListNode{
						8,
						nil,
					},
				},
				&ListNode{
					1,
					nil,
				},
			},
			&ListNode{
				0,
				&ListNode{
					9,
					nil,
				},
			},
		},
		{
			"[Test Case 5]",
			args{
				&ListNode{
					9,
					&ListNode{
						9,
						nil,
					},
				},
				&ListNode{
					1,
					nil,
				},
			},
			&ListNode{
				0,
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
			"[Test Case 6]",
			args{
				&ListNode{
					1,
					nil,
				},
				&ListNode{
					9,
					&ListNode{
						9,
						nil,
					},
				},
			},
			&ListNode{
				0,
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
			"[Test Case 7]",
			args{
				&ListNode{
					1,
					nil,
				},
				&ListNode{
					9,
					&ListNode{
						9,
						&ListNode{
							9,
							nil,
						},
					},
				},
			},
			&ListNode{
				0,
				&ListNode{
					0,
					&ListNode{
						0,
						&ListNode{
							1,
							nil,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
