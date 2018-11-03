package leetcode

import "testing"

func Test_minEatingSpeed(t *testing.T) {
	type args struct {
		piles []int
		H     int
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
					3, 6, 7, 11,
				},
				8,
			},
			4,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					30, 11, 23, 4, 20,
				},
				5,
			},
			30,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					30, 11, 23, 4, 20,
				},
				6,
			},
			23,
		},
		{
			"[Test Case 4]",
			args{
				[]int{
					3, 6, 7, 11,
				},
				8,
			},
			4,
		},
		{
			"[Test Case 5]",
			args{
				[]int{
					312884470,
				},
				968709470,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEatingSpeed(tt.args.piles, tt.args.H); got != tt.want {
				t.Errorf("minEatingSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
