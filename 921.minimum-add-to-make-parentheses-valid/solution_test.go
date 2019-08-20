package leetcode

import "testing"

func Test_minAddToMakeValid(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				"())",
			},
			1,
		},
		{
			"[Test Case 2]",
			args{
				"(((",
			},
			3,
		},
		{
			"[Test Case 3]",
			args{
				"()",
			},
			0,
		},
		{
			"[Test Case 4]",
			args{
				"()))((",
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAddToMakeValid(tt.args.S); got != tt.want {
				t.Errorf("minAddToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
