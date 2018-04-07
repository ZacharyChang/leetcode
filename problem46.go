// Given a collection of distinct numbers, return all possible permutations.

// For example,
// [1,2,3] have the following permutations:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]

package leetcode

func permute(nums []int) [][]int {
	var result [][]int
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{{
			nums[0],
		}}
	}
	for i, num := range nums {
		// make a copy, otherwise the nums may change after nums[:i]
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		excludeNums := append(numsCopy[:i], numsCopy[i+1:]...)

		temp := permute(excludeNums)
		for _, tempNum := range temp {
			result = append(result, append([]int{num}, tempNum...))
		}
	}
	return result
}
