package leetcode

import (
	"reflect"
	"strings"
	"testing"
)

func Test_findDuplicate(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"[Test Case 1]",
			args{
				[]string{
					"root/a 1.txt(abcd) 2.txt(efgh)",
					"root/c 3.txt(abcd)",
					"root/c/d 4.txt(efgh)",
					"root 4.txt(efgh)",
				},
			},
			[][]string{
				{
					"root/a/2.txt",
					"root/c/d/4.txt",
					"root/4.txt",
				},
				{
					"root/a/1.txt",
					"root/c/3.txt",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.paths); !equal(got, tt.want) {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a [][]string, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	aMap := make(map[string]bool, 0)
	bMap := make(map[string]bool, 0)
	for i := range a {
		aMap[strings.Join(a[i], ",")] = true
		bMap[strings.Join(b[i], ",")] = true
	}
	return reflect.DeepEqual(aMap, bMap)
}
