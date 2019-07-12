package leetcode

type WordFilter struct {
	prefixMap map[string][]int
	suffixMap map[string][]int
	length    int
}

func Constructor(words []string) WordFilter {
	res := WordFilter{
		prefixMap: make(map[string][]int, 0),
		suffixMap: make(map[string][]int, 0),
		length:    len(words),
	}
	for index, word := range words {
		for j := 1; j <= len(word); j++ {
			res.prefixMap[word[:j]] = append(res.prefixMap[word[:j]], index)
			res.suffixMap[word[len(word)-j:]] = append(res.suffixMap[word[len(word)-j:]], index)
		}
	}
	return res
}

func (this *WordFilter) F(prefix string, suffix string) int {
	if prefix == "" && suffix == "" {
		return this.length - 1
	}
	pre := this.prefixMap[prefix]
	suf := this.suffixMap[suffix]
	if prefix == "" {
		return findMaxWeight(suf, suf)
	}
	if suffix == "" {
		return findMaxWeight(pre, pre)
	}
	return findMaxWeight(pre, suf)
}

func findMaxWeight(prefixWeights []int, suffixWeights []int) int {
	for i := len(prefixWeights) - 1; i >= 0; i-- {
		for j := len(suffixWeights) - 1; j >= 0; j-- {
			if prefixWeights[i] == suffixWeights[j] {
				return prefixWeights[i]
			} else if prefixWeights[i] > suffixWeights[j] {
				break
			}
		}
	}
	return -1
}
