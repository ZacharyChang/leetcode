package leetcode

import (
	"reflect"
	"testing"
)

func Test_prisonAfterNDays(t *testing.T) {
	type args struct {
		cells []int
		N     int
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
					0, 1, 0, 1, 1, 0, 0, 1,
				},
				7,
			},
			[]int{
				0, 0, 1, 1, 0, 0, 0, 0,
			},
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 0, 0, 1, 0, 0, 1, 0,
				},
				1000000000,
			},
			[]int{
				0, 0, 1, 1, 1, 1, 1, 0,
			},
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					1, 0, 0, 1, 0, 0, 0, 1,
				},
				826,
			},
			[]int{
				0, 1, 1, 0, 1, 1, 1, 0,
			},
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					1, 0, 0, 1, 0, 0, 0, 1,
				},
				0,
			},
			[]int{
				1, 0, 0, 1, 0, 0, 0, 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prisonAfterNDays(tt.args.cells, tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prisonAfterNDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
