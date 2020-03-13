package leetcode

import "testing"

func Test_licenseKeyFormatting(t *testing.T) {
	type args struct {
		S string
		K int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"[Test Case 1]",
			args{
				"5F3Z-2e-9-w",
				4,
			},
			"5F3Z-2E9W",
		},
		{
			"[Test Case 2]",
			args{
				"2-5g-3-J",
				2,
			},
			"2-5G-3J",
		},
		{
			"[Test Case 3]",
			args{
				"2-4A0r7-4k",
				4,
			},
			"24A0-R74K",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := licenseKeyFormatting(tt.args.S, tt.args.K); got != tt.want {
				t.Errorf("licenseKeyFormatting() = %v, want %v", got, tt.want)
			}
		})
	}
}
