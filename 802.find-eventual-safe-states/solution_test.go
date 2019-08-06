package leetcode

import (
	"reflect"
	"testing"
)

func Test_eventualSafeNodes(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"[Test Case 1]",
			args{
				[][]int{
					{1, 2},
					{2, 3},
					{5},
					{0},
					{5},
					{},
					{},
				},
			},
			[]int{
				2, 4, 5, 6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eventualSafeNodes(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventualSafeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
