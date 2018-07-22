package leetcode

import (
	"reflect"
	"testing"
)

func Test_countBits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"[Test Case 1]",
			args{
				5,
			},
			[]int{
				0, 1, 1, 2, 1, 2,
			},
		},
		{
			"[Test Case 2]",
			args{
				10,
			},
			[]int{
				0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2,
			},
		},
		{
			"[Test Case 3]",
			args{
				0,
			},
			[]int{
				0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
