package leetcode

import "testing"

func Test_minSteps(t *testing.T) {
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
			3,
		},
		{
			"[Test Case 2]",
			args{
				6,
			},
			5,
		},
		{
			"[Test Case 3]",
			args{
				10,
			},
			7,
		},
		{
			"[Test Case 4]",
			args{
				1,
			},
			0,
		},
		{
			"[Test Case 5]",
			args{
				9,
			},
			6,
		},
		{
			"[Test Case 6]",
			args{
				7,
			},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSteps(tt.args.n); got != tt.want {
				t.Errorf("minSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
