package leetcode

import "testing"

func Test_intToRoman_2(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"[Test Case 1]",
			args{
				3999,
			},
			"MMMCMXCIX",
		},
		{
			"[Test Case 2]",
			args{
				5,
			},
			"V",
		},
		{
			"[Test Case 3]",
			args{
				1992,
			},
			"MCMXCII",
		},
		{
			"[Test Case 4]",
			args{
				1647,
			},
			"MDCXLVII",
		},
		{
			"[Test Case 5]",
			args{
				7,
			},
			"VII",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman_2(tt.args.num); got != tt.want {
				t.Errorf("intToRoman_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
