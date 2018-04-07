// Given an array of strings, group anagrams together.

// For example, given: ["eat", "tea", "tan", "ate", "nat", "bat"],
// Return:

// [
//   ["ate", "eat","tea"],
//   ["nat","tan"],
//   ["bat"]
// ]
// Note: All inputs will be in lower-case.
package leetcode

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)
	var result [][]string
	for _, v := range strs {
		str := strings.Split(v, "")
		sort.Strings(str)
		strMap[strings.Join(str, "")] = append(strMap[strings.Join(str, "")], v)
	}
	for _, v := range strMap {
		result = append(result, v)
	}
	return result
}
