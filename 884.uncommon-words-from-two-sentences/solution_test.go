package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func Test_uncommonFromSentences(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"[Test Case 1]",
			args{
				"this apple is sweet",
				"this apple is sour",
			},
			[]string{
				"sweet",
				"sour",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uncommonFromSentences(tt.args.A, tt.args.B); !equal(got, tt.want) {
				t.Errorf("uncommonFromSentences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)
	return reflect.DeepEqual(a, b)
}
