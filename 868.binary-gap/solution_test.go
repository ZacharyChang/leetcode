package leetcode

import "testing"

func Test_binaryGap(t *testing.T) {
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
				22,
			},
			2,
		},
		{
			"[Test Case 2]",
			args{
				5,
			},
			2,
		},
		{
			"[Test Case 3]",
			args{
				1234,
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryGap(tt.args.N); got != tt.want {
				t.Errorf("binaryGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
