package leetcode

import "testing"

func Test_complexNumberMultiply(t *testing.T) {
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
				"1+1i",
				"1+1i",
			},
			"0+2i",
		},
		{
			"[Test Case 2]",
			args{
				"1+-1i",
				"1+-1i",
			},
			"0+-2i",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := complexNumberMultiply(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("complexNumberMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
