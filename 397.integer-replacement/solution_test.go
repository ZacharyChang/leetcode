package leetcode

import "testing"

func Test_integerReplacement(t *testing.T) {
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
				8,
			},
			3,
		},
		{
			"[Test Case 2]",
			args{
				7,
			},
			4,
		},
		{
			"[Test Case 3]",
			args{
				15,
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerReplacement(tt.args.n); got != tt.want {
				t.Errorf("integerReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
