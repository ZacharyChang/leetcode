package leetcode

// slide window
func findAnagrams(s string, p string) []int {
	left, right := 0, 0
	// 差异度
	numDiff := len(p)
	res := make([]int, 0)

	charArray := make([]int, 26)
	for _, v := range p {
		charArray[v-'a']++
	}

	for ; right < len(s); right++ {
		charArray[s[right]-'a']--
		// 右窗口的元素为待匹配的元素时，差异度减1
		if charArray[s[right]-'a'] >= 0 {
			numDiff--
		}
		if right-left == len(p)-1 {
			// 差异度为0时，表示当前滑动窗口匹配
			if numDiff == 0 {
				res = append(res, left)
			}
			// 滑动左窗口指针
			// 左窗口的元素为已匹配元素时，差异度加1
			if charArray[s[left]-'a'] >= 0 {
				numDiff++
			}
			charArray[s[left]-'a']++
			// 左窗口向右滑动
			left++
		}
	}
	return res
}
