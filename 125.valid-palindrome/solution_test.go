package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				"A man, a plan, a canal: Panama",
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				"0P",
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				"`l;`` 1o1 ??;l`",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
