// Implement strStr().

package leetcode

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				"hello",
				"ll",
			},
			2,
		},
		{
			"[Test Case 2]",
			args{
				"aaaaa",
				"bba",
			},
			-1,
		},
		{
			"[Test Case 3]",
			args{
				"mississippi",
				"pi",
			},
			9,
		},
		{
			"[Test Case 3]",
			args{
				"",
				"",
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
