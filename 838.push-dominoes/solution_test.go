package leetcode

import "testing"

func Test_pushDominoes(t *testing.T) {
	type args struct {
		dominoes string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"[Test Case 1]",
			args{
				".L.R...LR..L..",
			},
			"LL.RR.LLRRLL..",
		},
		{
			"[Test Case 2]",
			args{
				"RR.L",
			},
			"RR.L",
		},
		{
			"[Test Case 3]",
			args{
				".L.R.",
			},
			"LL.RR",
		},
		{
			"[Test Case 4",
			args{
				"..R..",
			},
			"..RRR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pushDominoes(tt.args.dominoes); got != tt.want {
				t.Errorf("pushDominoes() = %v, want %v", got, tt.want)
			}
		})
	}
}
