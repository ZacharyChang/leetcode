package leetcode

import "testing"

func Test_canConstruct_loop(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				"aa",
				"aab",
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				"aa",
				"ab",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct_loop(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct_loop() = %v, want %v", got, tt.want)
			}
		})
	}
}
