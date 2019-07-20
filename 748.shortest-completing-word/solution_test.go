package leetcode

import "testing"

func Test_shortestCompletingWord(t *testing.T) {
	type args struct {
		licensePlate string
		words        []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"[Test Case 1]",
			args{
				"1s3 PSt",
				[]string{
					"step",
					"steps",
					"stripe",
					"stepple",
				},
			},
			"steps",
		},
		{
			"[Test Case 2]",
			args{
				"1s3 456",
				[]string{
					"looks",
					"pest",
					"stew",
					"show",
				},
			},
			"pest",
		},
		{
			"[Test Case 3]",
			args{
				"re69865",
				[]string{
					"population", "crime", "kid", "pressure", "store", "any", "relate", "will", "death", "when",
				},
			},
			"crime",
		},
		{
			"[Test Case 4]",
			args{
				"re65b",
				[]string{
					"population", "crime", "kid", "pressure", "store", "any", "relate", "will", "death", "when",
				},
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestCompletingWord(tt.args.licensePlate, tt.args.words); got != tt.want {
				t.Errorf("shortestCompletingWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
