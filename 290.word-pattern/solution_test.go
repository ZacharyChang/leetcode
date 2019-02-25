package leetcode

import "testing"

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		str     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				"abba",
				"dog cat cat dog",
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				"abba",
				"dog cat cat fish",
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				"jquery",
				"jquery",
			},
			false,
		},
		{
			"[Test Case 4]",
			args{
				"abba",
				"dog dog dog dog",
			},
			false,
		},
		{
			"[Test Case 5]",
			args{
				"e",
				"eureka",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.str); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
