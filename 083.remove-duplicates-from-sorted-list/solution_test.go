package leetcode

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"[Tets Case 1]",
			args{
				&ListNode{
					1,
					&ListNode{
						1,
						&ListNode{
							2,
							nil,
						},
					},
				},
			},
			&ListNode{
				1,
				&ListNode{
					2,
					nil,
				},
			},
		},
		{
			"[Tets Case 2]",
			args{
				&ListNode{
					1,
					&ListNode{
						1,
						&ListNode{
							2,
							&ListNode{
								3,
								&ListNode{
									3,
									nil,
								},
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
						nil,
					},
				},
			},
		},
		{
			"[Test Case 3]",
			args{
				nil,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
