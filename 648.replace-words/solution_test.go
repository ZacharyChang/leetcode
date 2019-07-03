package leetcode

import "testing"

func Test_replaceWords(t *testing.T) {
	type args struct {
		dict     []string
		sentence string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"[Test Case 1]",
			args{
				[]string{
					"cat",
					"bat",
					"rat",
				},
				"the cattle was rattled by the battery",
			},
			"the cat was rat by the bat",
		},
		{
			"[Test Case 2]",
			args{
				[]string{
					"cat",
					"bat",
					"rat",
				},
				"",
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceWords(tt.args.dict, tt.args.sentence); got != tt.want {
				t.Errorf("replaceWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
