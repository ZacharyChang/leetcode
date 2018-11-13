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
			"[Test Case 1]",
			args{
				&ListNode{
					1,
					&ListNode{
						2,
						&ListNode{
							3,
							&ListNode{
								3,
								&ListNode{
									4,
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
					},
				},
			},
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						5,
						nil,
					},
				},
			},
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
