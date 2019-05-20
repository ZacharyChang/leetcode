package leetcode

import "testing"

func Test_longestPalindrome_2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"[Test Case 1]",
			args{
				"babad",
			},
			"aba",
		},
		{
			"[Test Case 2]",
			args{
				"",
			},
			"",
		},
		{
			"[Test Case 3]",
			args{
				"a",
			},
			"a",
		},
		{
			"[Test Case 4",
			args{
				"civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth",
			},
			"ranynar",
		},
		{
			"[Test Case 5]",
			args{
				"cbbd",
			},
			"bb",
		},
		{
			"[Test Case 6]",
			args{
				"ac",
			},
			"a",
		},
		{
			"[Test Case 7]",
			args{
				"aaaa",
			},
			"aaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome_2(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
