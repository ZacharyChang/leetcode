package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func Test_accountsMerge(t *testing.T) {
	type args struct {
		accounts [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"[Test Case 1]",
			args{
				[][]string{
					{"John", "johnsmith@mail.com", "john00@mail.com"},
					{"John", "johnnybravo@mail.com"},
					{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
					{"Mary", "mary@mail.com"},
				},
			},
			[][]string{
				{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"Mary", "mary@mail.com"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := accountsMerge(tt.args.accounts)
			w := tt.want
			sort.Sort(emailList(got))
			sort.Sort(emailList(w))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("accountsMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

type emailList [][]string

func (a emailList) Len() int { return len(a) }
func (a emailList) Less(i, j int) bool {
	if len(a[i]) < len(a[j]) {
		return true
	} else if len(a[i]) == len(a[j]) {
		for m := 0; m < len(a[i]); m++ {
			if a[i][m] < a[j][m] {
				return true
			} else if a[i][m] > a[j][m] {
				return false
			}
		}
	}
	return false
}
func (a emailList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
