package leetcode

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 {
		return nums
	}
	count := len(nums) * len(nums[0])
	if count <= (r-1)*c || count > r*c {
		return nums
	}
	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			pos := i*len(nums[0]) + j
			res[pos / c][pos % c] = nums[i][j]
		}
	}
	return res
}
