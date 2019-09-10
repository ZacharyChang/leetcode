package leetcode

import (
	"reflect"
	"testing"
)

func Test_addToArrayForm(t *testing.T) {
	type args struct {
		A []int
		K int
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
					1, 2, 0, 0,
				},
				34,
			},
			[]int{
				1, 2, 3, 4,
			},
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 2, 3, 4,
				},
				8766,
			},
			[]int{
				1, 0, 0, 0, 0,
			},
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					9, 5,
				},
				6,
			},
			[]int{
				1, 0, 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addToArrayForm(tt.args.A, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addToArrayForm() = %v, want %v", got, tt.want)
			}
		})
	}
}
