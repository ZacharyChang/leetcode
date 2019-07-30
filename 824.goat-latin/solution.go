package leetcode

import "strings"

func toGoatLatin(S string) string {
	arr := strings.Split(S, " ")
	for i, v := range arr {
		arr[i] = toGoatLatinWord(v, i+1)
	}
	return strings.Join(arr, " ")
}

func toGoatLatinWord(w string, count int) string {
	begin := w[0]
	res := ""
	if begin == 'a' || begin == 'e' || begin == 'i' || begin == 'o' || begin == 'u' || begin == 'A' || begin == 'E' || begin == 'I' || begin == 'O' || begin == 'U' {
		res = w + "ma"
	} else {
		res = w[1:] + w[:1] + "ma"
	}
	for i := 0; i < count; i++ {
		res += "a"
	}
	return res
}
