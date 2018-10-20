package leetcode

import "testing"

func Test_findRadius(t *testing.T) {
	type args struct {
		houses  []int
		heaters []int
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
				[]int{
					2,
				},
			},
			1,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 2, 3, 4,
				},
				[]int{
					1, 4,
				},
			},
			1,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					1, 5,
				},
				[]int{
					2,
				},
			},
			3,
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					282475249, 622650073, 984943658, 144108930, 470211272, 101027544, 457850878, 458777923,
				},
				[]int{
					823564440, 115438165, 784484492, 74243042, 114807987, 137522503, 441282327, 16531729, 823378840, 143542612,
				},
			},
			161834419,
		},
		{
			"[Test Case 5]",
			args{
				[]int{},
				[]int{
					1,
				},
			},
			0,
		},
		{
			"[Test Case 6]",
			args{
				[]int{
					1, 2, 3, 4, 5, 6,
				},
				[]int{
					1, 3, 5,
				},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRadius(tt.args.houses, tt.args.heaters); got != tt.want {
				t.Errorf("findRadius() = %v, want %v", got, tt.want)
			}
		})
	}
}
