package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func Test_commonChars(t *testing.T) {
	type args struct {
		A []string
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
					"bella", "label", "roller",
				},
			},
			[]string{
				"e", "l", "l",
			},
		},
		{
			"[Test Case 2]",
			args{
				[]string{
					"cool", "lock", "cook",
				},
			},
			[]string{
				"c", "o",
			},
		},
		{
			"[Test Case 3]",
			args{
				[]string{
					"cool", "lock", "co",
				},
			},
			[]string{
				"c", "o",
			},
		},
		{
			"[Test Case 4]",
			args{
				[]string{},
			},
			[]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := commonChars(tt.args.A)
			sort.Strings(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
