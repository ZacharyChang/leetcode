package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want twoDimensionalArray
	}{
		{
			"[Test Case 1]",
			args{
				[]int{
					2, 3, 6, 7,
				},
				7,
			},
			[][]int{
				{
					7,
				},
				{
					2, 2, 3,
				},
			},
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					3, 6, 7,
				},
				7,
			},
			[][]int{
				{
					7,
				},
			},
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					2, 3, 5,
				},
				8,
			},
			[][]int{
				{
					2, 2, 2, 2,
				},
				{
					2, 3, 3,
				},
				{
					3, 5,
				},
			},
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					3, 5,
				},
				6,
			},
			[][]int{
				{
					3, 3,
				},
			},
		},
		{
			"[Test Case 5]",
			args{
				[]int{
					5,
				},
				6,
			},
			[][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !equal(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

type twoDimensionalArray [][]int

func (c twoDimensionalArray) Len() int {
	return len(c)
}

func (c twoDimensionalArray) Less(i, j int) bool {
	return len(c[i]) < len(c[j])
}

func (c twoDimensionalArray) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func equal(a twoDimensionalArray, b twoDimensionalArray) bool {
	sort.Sort(a)
	sort.Sort(b)
	return reflect.DeepEqual(a, b)
}
