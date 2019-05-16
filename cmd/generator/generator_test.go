package main

import "testing"

func Test_camel2Snake(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"[Test Case 1]",
			args{
				"Array",
			},
			"array",
		},
		{
			"[Test Case 2]",
			args{
				"BinarySearch",
			},
			"binary-search",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := camel2Snake(tt.args.str); got != tt.want {
				t.Errorf("camel2Snake() = %v, want %v", got, tt.want)
			}
		})
	}
}
