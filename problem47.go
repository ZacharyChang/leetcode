// Given a collection of numbers that might contain duplicates, return all possible unique permutations.

// For example,
// [1,1,2] have the following unique permutations:
// [
//   [1,1,2],
//   [1,2,1],
//   [2,1,1]
// ]
package leetcode

import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	var result [][]int
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{{
			nums[0],
		}}
	}
	sort.Ints(nums)
	for i, num := range nums {
		// skip the repeat element(already sorted)
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// make a copy, otherwise the nums may change after nums[:i]
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		excludeNums := append(numsCopy[:i], numsCopy[i+1:]...)

		temp := permuteUnique(excludeNums)
		for _, tempNum := range temp {
			result = append(result, append([]int{num}, tempNum...))
		}
	}
	return result
}
