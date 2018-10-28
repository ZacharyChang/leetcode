package leetcode

import "testing"

func Test_convertToTitle(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {
		// 	"[Test Case 1]",
		// 	args{
		// 		1,
		// 	},
		// 	"A",
		// },
		{
			"[Test Case 2]",
			args{
				701,
			},
			"ZY",
		},
		{
			"[Test Case 3]",
			args{
				26,
			},
			"Z",
		},
		{
			"[Test Case 4]",
			args{
				27,
			},
			"AA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args.n); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
