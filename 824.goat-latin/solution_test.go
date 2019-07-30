package leetcode

import "testing"

func Test_toGoatLatin(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"[Test Case 1]",
			args{
				"I speak Goat Latin",
			},
			"Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		},
		{
			"[Test Case 2]",
			args{
				"The quick brown fox jumped over the lazy dog",
			},
			"heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
		{
			"[Test Case 3]",
			args{
				"Each word consists of lowercase and uppercase letters only",
			},
			"Eachmaa ordwmaaa onsistscmaaaa ofmaaaaa owercaselmaaaaaa andmaaaaaaa uppercasemaaaaaaaa etterslmaaaaaaaaa onlymaaaaaaaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toGoatLatin(tt.args.S); got != tt.want {
				t.Errorf("toGoatLatin() = %v, want %v", got, tt.want)
			}
		})
	}
}
