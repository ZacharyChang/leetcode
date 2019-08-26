package leetcode

import "testing"

func Test_isLongPressedName(t *testing.T) {
	type args struct {
		name  string
		typed string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				"alex",
				"aaleex",
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				"saeed",
				"ssaaaedd",
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				"leelee",
				"lleeelee",
			},
			true,
		},
		{
			"[Test Case 4]",
			args{
				"laiden",
				"laiden",
			},
			true,
		},
		{
			"[Test Case 5]",
			args{
				"laiden",
				"laiien",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLongPressedName(tt.args.name, tt.args.typed); got != tt.want {
				t.Errorf("isLongPressedName() = %v, want %v", got, tt.want)
			}
		})
	}
}
