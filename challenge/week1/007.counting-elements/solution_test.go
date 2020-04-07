package leetcode

import "testing"

func Test_countElements(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				[]int{
					1, 2, 3,
				},
			},
			2,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 1, 3, 3, 5, 5, 7, 7,
				},
			},
			0,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					1, 3, 2, 3, 5, 0,
				},
			},
			3,
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					1, 1, 2, 2,
				},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countElements(tt.args.arr); got != tt.want {
				t.Errorf("countElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
