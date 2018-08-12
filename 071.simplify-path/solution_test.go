package leetcode

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"[Test Case 1]",
			args{
				"/a/./b/../../c/",
			},
			"/c",
		},
		{
			"[Test Case 2]",
			args{
				"/home",
			},
			"/home",
		},
		{
			"[Test Case 3]",
			args{
				"/",
			},
			"/",
		},
		{
			"[Test Case 4]",
			args{
				"/home/../../..",
			},
			"/",
		},
		{
			"[Test Case 5]",
			args{
				"///",
			},
			"/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
