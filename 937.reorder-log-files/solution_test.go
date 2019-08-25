package leetcode

import (
	"reflect"
	"testing"
)

func Test_reorderLogFiles(t *testing.T) {
	type args struct {
		logs []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"[Test Case 1]",
			args{
				[]string{
					"a1 9 2 3 1", "g1 act car",
					"zo4 4 7", "ab1 off key dog",
					"a8 act zoo",
				},
			},
			[]string{
				"g1 act car",
				"a8 act zoo",
				"ab1 off key dog",
				"a1 9 2 3 1",
				"zo4 4 7",
			},
		},
		{
			"[Test Case 2]",
			args{
				[]string{
					"a1 9 2 3 1",
					"g1 act car",
					"zo4 4 7",
					"ab1 off key dog",
					"a8 act zoo",
					"a2 act car",
				},
			},
			[]string{
				"a2 act car",
				"g1 act car",
				"a8 act zoo",
				"ab1 off key dog",
				"a1 9 2 3 1",
				"zo4 4 7",
			},
		},
		{
			"[Test Case 3]",
			args{
				[]string{
					"wpylev6cnqv8 447241070789889321113517804297550370", "2syod 60098540876848105552318 69698830167476212 2", "iuw2x1r qmxealfvosqgkv yunonsq nxcuwudndrpsroty h", "vclnqwkdr7h 5515 892770977116949342793820104705 3", "5y08u4f5ba swixvlwipfhtxavvzrdyxtnronckklvh f kzd", "2k63p1p 542447297738584 22237063423417308275099706", "qrj 467859 382 451796621324556 12022 72631305 0429", "vopck4 huqziggmwvcsermnujnpplihttviwei lsrqmbw b n", "s 7257018672440110203152567646 961657508453405583", "94j 1800907 54116251858 19612167 218608 1 504204 4", "u34lvgmoh 631217074786612695089137448 5635620839 5", "dsrojn8aeojx 27159799084241651870 76594680 195 051", "kteuav 77685739 6366458436088787165747310 78 3849", "dg8gco0sc2 10811560194867165521681 718 42498 1 52", "gdg 900670532316533434070453812 9115641245822 122", "ytlmfg 658910166131 170942932 70238 0783568 64777", "gi6d2lg8 ekwbnzeqrrzijgexvpcgfnhrkfhtmegaqon hsa u", "y8zhzn jjtbwpfrbcsuj zmseejb vcsovstaewtgtj nbsnlj", "cp1qsk5 dstuzhik alqxnmztxnwdve simoynsfffwyacl nr", "a11zjdza15gi zuighjavkfidjjx kgmbriwxbjcsv shtfjz", "3dpx3f28wa1 dhe jb uatgwooxclfj w t qaahvyiy lthj", "9ymabvjk4xq waujeijoltuy heoekaqmggmpdkynngne sl x", "n3l09gzpppgc dfnfxeaskknllixe tvtbemewtkwd bfbhm l", "one 143245418564431590 555274555077126490673 23406", "2dlvtxq57 11 0399853766553806651446400571374174 7", "l0xsyrtf9foe jcsjyzbux hpxxwwsyxwjcdqbuzrxuvdf n n", "o ufjxgmiohhacgwhprzqklpbleggurqygvmyrqtiwwusaa fq", "gr jmhsaanfrndkvkrdepfqvnathkheq bjtvzacabyfch xw", "12hrfmpyxql 509513107446443470266800090 12 36792 5", "ei mfmrujazj hvcaeoejhbvsxlnbcofdparedjvuqoigbwv h", "158mo1 fxwvcxyaz gimthvk t tbkpxnyomitu i foi t i", "1mpnz91abn8 857526216042344 529 86 555850 074136 6", "gvf69aycgz vd wzshq vqqcoscdfgtclfpoqz kcnbk yqrta", "05tv1dyuuln 3010253552744498232332 86540056 5488 0", "1kjt2sp 76661129172454994454966100212569762 877775", "k8fq mhahouacluusiypbhmbxknagj nrenkpsijov tspqd s", "a 05783356043073570183098305205075240023467 24 63", "0c 821 1288302446431573478713998604686702 0584599", "e 02985850443721 356058 49996149758367 64432663 32", "1zayns7ifztd kwmwsthxzxvvctzoejspeobtavhzxzpot u n", "o0sh3 qn nqjaghnmkckhvuauuknqbuxwalgva nt gfhqm en", "094qnly wgkmupkjobuup gshx wcblufmjumyuahsx n ai k", "j69r2ugwa6 zuoywue chhwsfdprfygvliwzmohqgrxn ubwtm", "2mkuap uqfwog jqzrkoorsompgkdlql wpauhkzvig ftb l", "x 929 4356109428379557082235487428356570127401 832", "jns07q8 idnlfsaezcojuafbgmancqpegbzy q qwesz rmy n", "phk1cna 086 027760883273 64658492093523655560824 2", "jbemfs9l9bs0 8147538504741452659388775 5 32 180 09", "ac9cwb9 524689619771630155 8125241949139653850678", "9eke perwsfqykyslfmcwnovenuiy urstqeqaezuankek czq",
				},
			},
			[]string{
				"n3l09gzpppgc dfnfxeaskknllixe tvtbemewtkwd bfbhm l", "3dpx3f28wa1 dhe jb uatgwooxclfj w t qaahvyiy lthj", "cp1qsk5 dstuzhik alqxnmztxnwdve simoynsfffwyacl nr", "gi6d2lg8 ekwbnzeqrrzijgexvpcgfnhrkfhtmegaqon hsa u", "158mo1 fxwvcxyaz gimthvk t tbkpxnyomitu i foi t i", "vopck4 huqziggmwvcsermnujnpplihttviwei lsrqmbw b n", "jns07q8 idnlfsaezcojuafbgmancqpegbzy q qwesz rmy n", "l0xsyrtf9foe jcsjyzbux hpxxwwsyxwjcdqbuzrxuvdf n n", "y8zhzn jjtbwpfrbcsuj zmseejb vcsovstaewtgtj nbsnlj", "gr jmhsaanfrndkvkrdepfqvnathkheq bjtvzacabyfch xw", "1zayns7ifztd kwmwsthxzxvvctzoejspeobtavhzxzpot u n", "ei mfmrujazj hvcaeoejhbvsxlnbcofdparedjvuqoigbwv h", "k8fq mhahouacluusiypbhmbxknagj nrenkpsijov tspqd s", "9eke perwsfqykyslfmcwnovenuiy urstqeqaezuankek czq", "iuw2x1r qmxealfvosqgkv yunonsq nxcuwudndrpsroty h", "o0sh3 qn nqjaghnmkckhvuauuknqbuxwalgva nt gfhqm en", "5y08u4f5ba swixvlwipfhtxavvzrdyxtnronckklvh f kzd", "o ufjxgmiohhacgwhprzqklpbleggurqygvmyrqtiwwusaa fq", "2mkuap uqfwog jqzrkoorsompgkdlql wpauhkzvig ftb l", "gvf69aycgz vd wzshq vqqcoscdfgtclfpoqz kcnbk yqrta", "9ymabvjk4xq waujeijoltuy heoekaqmggmpdkynngne sl x", "094qnly wgkmupkjobuup gshx wcblufmjumyuahsx n ai k", "a11zjdza15gi zuighjavkfidjjx kgmbriwxbjcsv shtfjz", "j69r2ugwa6 zuoywue chhwsfdprfygvliwzmohqgrxn ubwtm", "wpylev6cnqv8 447241070789889321113517804297550370", "2syod 60098540876848105552318 69698830167476212 2", "vclnqwkdr7h 5515 892770977116949342793820104705 3", "2k63p1p 542447297738584 22237063423417308275099706", "qrj 467859 382 451796621324556 12022 72631305 0429", "s 7257018672440110203152567646 961657508453405583", "94j 1800907 54116251858 19612167 218608 1 504204 4", "u34lvgmoh 631217074786612695089137448 5635620839 5", "dsrojn8aeojx 27159799084241651870 76594680 195 051", "kteuav 77685739 6366458436088787165747310 78 3849", "dg8gco0sc2 10811560194867165521681 718 42498 1 52", "gdg 900670532316533434070453812 9115641245822 122", "ytlmfg 658910166131 170942932 70238 0783568 64777", "one 143245418564431590 555274555077126490673 23406", "2dlvtxq57 11 0399853766553806651446400571374174 7", "12hrfmpyxql 509513107446443470266800090 12 36792 5", "1mpnz91abn8 857526216042344 529 86 555850 074136 6", "05tv1dyuuln 3010253552744498232332 86540056 5488 0", "1kjt2sp 76661129172454994454966100212569762 877775", "a 05783356043073570183098305205075240023467 24 63", "0c 821 1288302446431573478713998604686702 0584599", "e 02985850443721 356058 49996149758367 64432663 32", "x 929 4356109428379557082235487428356570127401 832", "phk1cna 086 027760883273 64658492093523655560824 2", "jbemfs9l9bs0 8147538504741452659388775 5 32 180 09", "ac9cwb9 524689619771630155 8125241949139653850678",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderLogFiles(tt.args.logs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderLogFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
