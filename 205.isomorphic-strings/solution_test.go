package leetcode

import "testing"

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				"egg",
				"add",
			},
			true,
		},
		{
			"[Test Case 2]",
			args{
				"foo",
				"bar",
			},
			false,
		},
		{
			"[Test Case 3]",
			args{
				"ab",
				"aa",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
