package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	type args struct {
		T []int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args{
				[]int{
					73, 74, 75, 71, 69, 72, 76, 73,
				},
			},
			[]int{
				1, 1, 4, 2, 1, 1, 0, 0,
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("[Test Case %d]", i), func(t *testing.T) {
			if got := dailyTemperatures(tt.args.T); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}
