package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func Test_powerfulIntegers(t *testing.T) {
	type args struct {
		x     int
		y     int
		bound int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"[Test Case 1]",
			args{
				2, 3, 10,
			},
			[]int{
				2, 3, 4, 5, 7, 9, 10,
			},
		},
		{
			"[Test Case 2]",
			args{
				3, 5, 15,
			},
			[]int{
				2, 4, 6, 8, 10, 14,
			},
		},
		{
			"[Test Case 3]",
			args{
				2, 1, 10,
			},
			[]int{
				2, 3, 5, 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := powerfulIntegers(tt.args.x, tt.args.y, tt.args.bound)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("powerfulIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
