package leetcode

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	if target <= 0 {
		return res
	}
	// end condition
	if len(candidates) == 1 {
		if target < candidates[0] || target%candidates[0] > 0 {
			return res
		}
		array := make([]int, 0)
		for i := 0; i < target/candidates[0]; i++ {
			array = append(array, candidates[0])
		}
		res = append(res, array)
		return res
	}
	// direct get target from candidates[0]
	if target%candidates[0] == 0 {
		arr := make([]int, 0)
		for i := 0; i < target/candidates[0]; i++ {
			arr = append(arr, candidates[0])
		}
		res = append(res, arr)
	}
	// first get n*candidates[0] from first element
	// then get target-n*candidates[0] from candidates[1:]
	for i := 1; i*candidates[0] < target; i++ {
		for _, v := range combinationSum(candidates[1:], target-i*candidates[0]) {
			arr := make([]int, 0)
			for j := 0; j < i; j++ {
				arr = append(arr, candidates[0])
			}
			arr = append(arr, v...)
			res = append(res, arr)
		}
	}
	// direct get target from candidates[1:]
	res = append(res, combinationSum(candidates[1:], target)...)
	return res
}
