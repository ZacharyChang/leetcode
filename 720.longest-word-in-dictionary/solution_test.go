package leetcode

import "testing"

func Test_longestWord(t *testing.T) {
	type args struct {
		words []string
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
					"w", "wo", "wor", "worl", "world",
				},
			},
			"world",
		},
		{
			"[Test Case 2]",
			args{
				[]string{
					"a", "banana", "app", "appl", "ap", "apply", "apple",
				},
			},
			"apple",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestWord(tt.args.words); got != tt.want {
				t.Errorf("longestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
