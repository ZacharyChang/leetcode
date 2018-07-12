package leetcode

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

type NumMatrix struct {
	array [][]int // 记录原数组
	sum   [][]int // 记录数组和
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{}
	}
	sum := make([][]int, len(matrix))
	copy(sum, matrix)
	for i := 0; i < len(sum); i++ {
		for j := 0; j < len(sum[0]); j++ {
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
		}
	}
	return NumMatrix{
		array: matrix,
		sum:   sum,
	}
}

// Sum(ABCD) = Sum(OD) - Sum(OB) - Sum(OC) + Sum(OA)
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	result := this.sum[row2][col2]
	if col1 > 0 {
		result -= this.sum[row2][col1-1]
	}
	if row1 > 0 {
		result -= this.sum[row1-1][col2]
	}
	if col1 > 0 && row1 > 0 {
		result += this.sum[row1-1][col1-1]
	}
	return result
}
