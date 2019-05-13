package leetcode

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int, 0)
	for index, v := range nums {
		if numMap[target-v] > 0 {
			return []int{
				numMap[target-v] - 1, index,
			}
		}
		numMap[v] = index + 1
	}
	return []int{0, 1}
}
