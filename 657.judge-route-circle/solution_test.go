package leetcode

import "testing"

func Test_judgeCircle(t *testing.T) {
	type args struct {
		moves string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test case 1]",
			args{
				"UD",
			},
			true,
		},
		{
			"[Test case 2]",
			args{
				"LL",
			},
			false,
		},
		{
			"[Test case 3]",
			args{
				"DURDLDRRLL",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeCircle(tt.args.moves); got != tt.want {
				t.Errorf("judgeCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}
