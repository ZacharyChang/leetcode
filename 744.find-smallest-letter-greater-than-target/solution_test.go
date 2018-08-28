package leetcode

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			"[Test Case 1]",
			args{
				[]byte{
					'c', 'f', 'j',
				},
				'a',
			},
			'c',
		},
		{
			"[Test Case 2]",
			args{
				[]byte{
					'c', 'f', 'j',
				},
				'c',
			},
			'f',
		},
		{
			"[Test Case 3]",
			args{
				[]byte{
					'c', 'f', 'j',
				},
				'd',
			},
			'f',
		},
		{
			"[Test Case 4]",
			args{
				[]byte{
					'c', 'f', 'j',
				},
				'g',
			},
			'j',
		},
		{
			"[Test Case 5]",
			args{
				[]byte{
					'c', 'f', 'j',
				},
				'j',
			},
			'c',
		},
		{
			"[Test Case 6]",
			args{
				[]byte{
					'c', 'f', 'j',
				},
				'k',
			},
			'c',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
