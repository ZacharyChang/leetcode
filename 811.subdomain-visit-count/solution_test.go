package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func Test_subdomainVisits(t *testing.T) {
	type args struct {
		cpdomains []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"[Test Case 1]",
			args{
				[]string{
					"9001 discuss.leetcode.com",
				},
			},
			[]string{
				"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subdomainVisits(tt.args.cpdomains); !equal(got, tt.want) {
				t.Errorf("subdomainVisits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a []string, b []string) bool {
	sort.Strings(a)
	sort.Strings(b)
	return reflect.DeepEqual(a, b)
}
