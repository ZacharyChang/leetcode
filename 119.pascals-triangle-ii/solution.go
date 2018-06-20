package leetcode

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	result[0] = 1
	for i := 1; i < rowIndex+1; i++ {
		result[i] = 1
		// 重点：从后往前计算，依次叠加
		for j := i - 1; j > 0; j-- {
			result[j] += result[j-1]
		}
	}
	return result
}
