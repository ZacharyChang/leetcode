package leetcode

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"[Test Case 1]",
			args{
				[]string{
					"a",
					"bc",
					"d",
				},
			},
			"",
		},
		{
			"[Test Case 2]",
			args{
				[]string{
					"a",
					"ac",
					"ad",
				},
			},
			"a",
		},
		{
			"[Test Case 3]",
			args{
				[]string{},
			},
			"",
		},
		{
			"[Test Case 4]",
			args{
				[]string{
					"abc",
					"ab",
					"abde",
					"abc",
				},
			},
			"ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
