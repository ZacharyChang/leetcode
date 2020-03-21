package _19_bulb_switcher

import "testing"

func Test_bulbSwitch(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[Test Case 1]",
			args{
				3,
			},
			1,
		},
		{
			"[Test Case 2]",
			args{
				99,
			},
			9,
		},
		{
			"[Test Case 3]",
			args{
				99999999,
			},
			9999,
		},
		{
			"[Test Case 4]",
			args{
				0,
			},
			0,
		},
		{
			"[Test Case 5]",
			args{
				1,
			},
			1,
		},
		{
			"[Test Case 6]",
			args{
				2,
			},
			1,
		},
		{
			"[Test Case 7]",
			args{
				4,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bulbSwitch(tt.args.n); got != tt.want {
				t.Errorf("bulbSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}
