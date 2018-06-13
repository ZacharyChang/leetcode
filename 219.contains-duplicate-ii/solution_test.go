package leetcode

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[Test Case 1]",
			args{
				[]int{
					1, 2, 3, 1, 2, 3,
				},
				2,
			},
			false,
		},
		{
			"[Test Case 2]",
			args{
				[]int{
					1, 0, 1, 1,
				},
				1,
			},
			true,
		},
		{
			"[Test Case 3]",
			args{
				[]int{
					1, 2, 3, 1,
				},
				3,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
