package leetcode

import "testing"

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"[Test Case 1]",
			args{
				3,
			},
			"21",
		},
		{
			"[Test Case 2]",
			args{
				2,
			},
			"11",
		},
		{
			"[Test Case 3]",
			args{
				4,
			},
			"1211",
		},
		{
			"[Test Case 4]",
			args{
				5,
			},
			"111221",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.args.n); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
