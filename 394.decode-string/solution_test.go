package leetcode

import "testing"

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"[Test Case 1]",
			args{
				"3[a]2[bc]",
			},
			"aaabcbc",
		},
		{
			"[Test Case 2]",
			args{
				"3[a2[c]]",
			},
			"accaccacc",
		},
		{
			"[Test Case 3]",
			args{
				"a2[c]",
			},
			"acc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
