package leetcode

import "testing"

func Test_equationsPossible(t *testing.T) {
	type args struct {
		equations []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				[]string{
					"a==b",
					"b!=a",
				},
			},
			false,
		},
		{
			"[Test Case 2]",
			args{
				[]string{
					"b==a",
					"a==b",
				},
			},
			true,
		},
		{
			"[Test Case 3]",
			args{
				[]string{
					"a==b",
					"e==c",
					"b==c",
					"a!=e",
				},
			},
			false,
		},
		{
			"[Test Case 4]",
			args{
				[]string{
					"a==b",
					"b!=c",
					"c==a",
				},
			},
			false,
		},
		{
			"[Test Case 5]",
			args{
				[]string{
					"c==a",
					"b==d",
					"x!=z",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equationsPossible(tt.args.equations); got != tt.want {
				t.Errorf("equationsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
