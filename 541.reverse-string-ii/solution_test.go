package leetcode

import "testing"

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"[Test Case 1]",
			args{
				"abcdefg",
				2,
			},
			"bacdfeg",
		},
		{
			"[Test Case 2]",
			args{
				"abcdefg",
				3,
			},
			"cbadefg",
		},
		{
			"[Test Case 3]",
			args{
				"abcdefg",
				8,
			},
			"gfedcba",
		},
		{
			"[Test Case 4]",
			args{
				"hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl",
				39,
			},
			"fdcqkmxwholhytmhafpesaentdvxginrjlyqzyhehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqllgsqddebemjanqcqnfkjmi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
