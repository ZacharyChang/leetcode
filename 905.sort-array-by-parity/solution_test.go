package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParity(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"[Test Case 1]",
			args{
				[]int{
					3, 1, 2, 4,
				},
			},
			[]int{
				2, 4, 3, 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParity(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParity() = %v, want %v", got, tt.want)
			}
		})
	}
}
