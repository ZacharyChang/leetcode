package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				&ListNode{
					1,
					&ListNode{
						0,
						&ListNode{
							0,
							nil,
						},
					},
				},
			},
			false,
		},
		{
			"[Test Case 2]",
			args{
				&ListNode{
					0,
					&ListNode{
						0,
						nil,
					},
				},
			},
			true,
		},
		{
			"[Test Case 3]",
			args{
				&ListNode{
					1,
					&ListNode{
						0,
						nil,
					},
				},
			},
			false,
		},
		{
			"[Test Case 4]",
			args{
				&ListNode{
					1,
					nil,
				},
			},
			true,
		},
		{
			"[Test Case 5]",
			args{
				&ListNode{
					1,
					&ListNode{
						0,
						&ListNode{
							1,
							nil,
						},
					},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
