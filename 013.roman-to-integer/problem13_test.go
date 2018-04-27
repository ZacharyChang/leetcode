// Given a roman numeral, convert it to an integer.

// Input is guaranteed to be within the range from 1 to 3999.

package leetcode

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				"VI",
			},
			6,
		},
		{
			"[Test Case 2]",
			args{
				"DCXXI",
			},
			621,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
