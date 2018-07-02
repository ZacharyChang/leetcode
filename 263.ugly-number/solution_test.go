package leetcode

import "testing"

func Test_isUgly(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"{Test Case 1]",
			args{
				6,
			},
			true,
		},
		{
			"{Test Case 2]",
			args{
				14,
			},
			false,
		},
		{
			"{Test Case 3]",
			args{
				8,
			},
			true,
		},
		{
			"{Test Case 4]",
			args{
				0,
			},
			false,
		},
		{
			"{Test Case 5]",
			args{
				30,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUgly(tt.args.num); got != tt.want {
				t.Errorf("isUgly() = %v, want %v", got, tt.want)
			}
		})
	}
}
