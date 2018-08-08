package leetcode

import "testing"

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"[Test Case 1]",
			args{
				"11",
				"1",
			},
			"100",
		},
		{
			"[Test Case 2]",
			args{
				"11",
				"0",
			},
			"11",
		},
		{
			"[Test Case 3]",
			args{
				"1010",
				"1011",
			},
			"10101",
		},
		{
			"[Test Case 4]",
			args{
				"1",
				"111",
			},
			"1000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
