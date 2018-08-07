package leetcode

import "testing"

func Test_trailingZeroest(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				3,
			},
			0,
		},
		{
			"[Test Case 2]",
			args{
				5,
			},
			1,
		},
		{
			"[Test Case 3]",
			args{
				6,
			},
			1,
		},
		{
			"[Test Case 4]",
			args{
				1808548329,
			},
			452137076,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.args.n); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
