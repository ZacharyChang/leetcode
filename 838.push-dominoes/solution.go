package leetcode

func pushDominoes(dominoes string) string {
	result := make([]byte, len(dominoes))
	forces := make([]int, len(dominoes))
	tmp := 0
	// from left to right
	for i := range dominoes {
		if dominoes[i] == 'R' {
			tmp = len(dominoes)
		} else if dominoes[i] == '.' {
			// the force can't be less than zero
			tmp = max(tmp-1, 0)
		} else {
			tmp = 0
		}
		forces[i] += tmp
	}
	// from right to left
	for i := len(dominoes) - 1; i >= 0; i-- {
		if dominoes[i] == 'L' {
			tmp = len(dominoes)
		} else if dominoes[i] == '.' {
			tmp = max(tmp-1, 0)
		} else {
			tmp = 0
		}
		forces[i] -= tmp
	}
	for i, v := range forces {
		if v == 0 {
			result[i] = '.'
		} else if v > 0 {
			result[i] = 'R'
		} else {
			result[i] = 'L'
		}
	}
	return string(result)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
