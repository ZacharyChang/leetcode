package leetcode

import (
	"reflect"
	"testing"
)

func Test_imageSmoother(t *testing.T) {
	type args struct {
		M [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"[Test Case 1]",
			args{
				[][]int{
					{
						1, 1, 1,
					},
					{
						1, 0, 1,
					},
					{
						1, 1, 1,
					},
				},
			},
			[][]int{
				{
					0, 0, 0,
				},
				{
					0, 0, 0,
				},
				{
					0, 0, 0,
				},
			},
		},
		{
			"[Test Case 2]",
			args{
				[][]int{
					{
						1,
					},
				},
			},
			[][]int{
				{
					1,
				},
			},
		},
		{
			"[Test Case 3]",
			args{
				[][]int{
					{
						1, 0,
					},
					{
						1, 1,
					},
				},
			},
			[][]int{
				{
					0, 0,
				},
				{
					0, 0,
				},
			},
		},
		{
			"[Test Case 4]",
			args{
				[][]int{
					{
						2, 3,
					},
				},
			},
			[][]int{
				{
					2, 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := imageSmoother(tt.args.M); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageSmoother() = %v, want %v", got, tt.want)
			}
		})
	}
}
