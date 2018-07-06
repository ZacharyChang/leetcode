package leetcode

import "testing"

func TestNumArray_SumRange(t *testing.T) {
	type fields struct {
		array []int
		sum   []int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			"[Test Case 1]",
			fields{
				array: Constructor([]int{
					-2, 0, 3, -5, 2, -1,
				}).array,
				sum: Constructor([]int{
					-2, 0, 3, -5, 2, -1,
				}).sum,
			},
			args{
				0, 2,
			},
			1,
		},
		{
			"[Test Case 2]",
			fields{
				array: Constructor([]int{
					-2, 0, 3, -5, 2, -1,
				}).array,
				sum: Constructor([]int{
					-2, 0, 3, -5, 2, -1,
				}).sum,
			},
			args{
				2, 5,
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &NumArray{
				array: tt.fields.array,
				sum:   tt.fields.sum,
			}
			if got := this.SumRange(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("NumArray.SumRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
