package leetcode

import "testing"

func Test_fib(t *testing.T) {
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
				1,
			},
			1,
		},
		{
			"[Test Case 2]",
			args{
				3,
			},
			2,
		},
		{
			"[Test Case 3]",
			args{
				4,
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.N); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
