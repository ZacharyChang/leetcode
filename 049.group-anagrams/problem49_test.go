package leetcode

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"[Test Case 1]",
			args{
				[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			[][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !hasSameElement(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func hasSameElement(a [][]string, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		flag := false
		for j := 0; j < len(b); j++ {
			if reflect.DeepEqual(a[i], b[j]) {
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
}
