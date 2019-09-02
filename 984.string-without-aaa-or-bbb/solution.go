package leetcode

func strWithout3a3b(A int, B int) string {
	if A > B {
		return helper(A, B, true)
	} else if A < B {
		return helper(A, B, false)
	}
	return helper(A, B, true)
}

func helper(A int, B int, fromA bool) string {
	res := ""
	if A < 3 && B < 3 {
		if fromA {
			for A > 0 {
				res += "a"
				A--
			}
			for B > 0 {
				res += "b"
				B--
			}
		} else {
			for B > 0 {
				res += "b"
				B--
			}
			for A > 0 {
				res += "a"
				A--
			}
		}
		return res
	}
	if A > B {
		res += "aab" + helper(A-2, B-1, true)
	} else if A < B {
		res += "bba" + helper(A-1, B-2, false)
	} else {
		res += "ab" + helper(A-1, B-1, true)
	}
	return res
}
