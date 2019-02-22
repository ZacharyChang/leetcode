package leetcode

import (
	"reflect"
	"testing"
)

func Test_findAnagrams_2(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"[Test Case 1]",
			args{
				"cbaebabacd",
				"abc",
			},
			[]int{
				0, 6,
			},
		},
		{
			"[Test Case 2]",
			args{
				"abab",
				"ab",
			},
			[]int{
				0, 1, 2,
			},
		},
		{
			"[Test Case 3]",
			args{
				"af",
				"be",
			},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams_2(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
