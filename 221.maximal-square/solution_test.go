package leetcode

import "testing"

func Test_maximalSquare(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				[][]byte{
					{
						'1', '0', '1', '0', '0',
					},
					{
						'1', '0', '1', '1', '1',
					},
					{
						'1', '1', '1', '1', '1',
					},
					{
						'1', '0', '0', '1', '0',
					},
				},
			},
			4,
		},
		{
			"[Test Case 2]",
			args{
				[][]byte{
					{
						'1',
					},
				},
			},
			1,
		},
		{
			"[Test Case 3]",
			args{
				[][]byte{},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.args.matrix); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
