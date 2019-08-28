package leetcode

import "testing"

func Test_isAlienSorted(t *testing.T) {
	type args struct {
		words []string
		order string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				[]string{
					"hello",
					"leetcode",
				},
				"hlabcdefgijkmnopqrstuvwxyz",
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				[]string{
					"word",
					"world",
					"row",
				},
				"worldabcefghijkmnpqstuvxyz",
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				[]string{
					"apple",
					"app",
				},
				"abcdefghijklmnopqrstuvwxyz",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlienSorted(tt.args.words, tt.args.order); got != tt.want {
				t.Errorf("isAlienSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
