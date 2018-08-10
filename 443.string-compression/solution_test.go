package leetcode

import "testing"

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				[]byte{
					'a', 'a', 'b', 'b', 'c', 'c', 'c',
				},
			},
			6,
		},
		{
			"[Test Case 2]",
			args{
				[]byte{
					'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b',
				},
			},
			4,
		},
		{
			"[Test Case 3]",
			args{
				[]byte{
					'a',
				},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.chars); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
		})
	}
}
