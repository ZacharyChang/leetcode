package leetcode

import "testing"

func Test_checkRecord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				"PPALLP",
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				"PPALLL",
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				"LALL",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkRecord(tt.args.s); got != tt.want {
				t.Errorf("checkRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}
