package leetcode

import "testing"

func Test_countPrimes(t *testing.T) {
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
				10,
			},
			4,
		},
		{
			"[Test Case 2]",
			args{
				499979,
			},
			41537,
		},
		{
			"[Test Case 3]",
			args{
				5,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimes(tt.args.n); got != tt.want {
				t.Errorf("countPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
