package leetcode

type NumArray struct {
	array []int
	sum   []int
}

func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums))
	if len(nums) > 0 {
		sum[0] = nums[0]
		for i := 1; i < len(nums); i++ {
			sum[i] = sum[i-1] + nums[i]
		}
	}
	return NumArray{
		array: nums,
		sum:   sum,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.sum[j]
	}
	return this.sum[j] - this.sum[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
