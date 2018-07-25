package leetcode

import "testing"

func Test_reverseString_2(t *testing.T) {
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
				"hello",
			},
			"olleh",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseString_2pointers(tt.args.s); got != tt.want {
				t.Errorf("reverseString_2pointers() = %v, want %v", got, tt.want)
			}
		})
	}
}
