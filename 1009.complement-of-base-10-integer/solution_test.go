package leetcode

import "testing"

func Test_bitwiseComplement(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				5,
			},
			2,
		},
		{
			"[Test Case 2]",
			args{
				7,
			},
			0,
		},
		{
			"[Test Case 3]",
			args{
				10,
			},
			5,
		},
		{
			"[Test Case 5]",
			args{
				0,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitwiseComplement(tt.args.N); got != tt.want {
				t.Errorf("bitwiseComplement() = %v, want %v", got, tt.want)
			}
		})
	}
}
