package leetcode

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		S string
		T string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				"ab#c",
				"ad#c",
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				"ab##",
				"c#d#",
			},
			true,
		},
		{
			"[Test Case 3]",
			args{
				"a##c",
				"#a#c",
			},
			true,
		},
		{
			"[Test Case 4]",
			args{
				"a#c",
				"b",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
