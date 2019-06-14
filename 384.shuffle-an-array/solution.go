package leetcode

import "math/rand"

// TODO: Fisher-Yates Algorithm
type Solution struct {
	Original     []int
	ShuffleArray []int
}

func Constructor(nums []int) Solution {
	return Solution{
		Original:     nums,
		ShuffleArray: nums,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	this.ShuffleArray = this.Original
	return this.ShuffleArray
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	arrCopy := make([]int, len(this.ShuffleArray))
	copy(arrCopy, this.Original)
	res := make([]int, 0)
	for len(arrCopy) > 0 {
		idx := rand.Intn(len(arrCopy))
		res = append(res, arrCopy[idx])
		arrCopy = append(arrCopy[:idx], arrCopy[idx+1:]...)
	}
	this.ShuffleArray = res
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
