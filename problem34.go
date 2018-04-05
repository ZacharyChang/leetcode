// Given an array of integers sorted in ascending order, find the starting and ending position of a given target value.

// Your algorithm's runtime complexity must be in the order of O(log n).

// If the target is not found in the array, return [-1, -1].

// For example,
// Given [5, 7, 7, 8, 8, 10] and target value 8,
// return [3, 4].

package leetcode

func searchRange(nums []int, target int) []int {
	return []int{findLeftIndex(nums, target), findRightIndex(nums, target)}
}

func findLeftIndex(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		middle := (low + high) / 2
		if nums[middle] > target { // 中间的元素比目标大时，查找左半边
			high = middle - 1
		} else if nums[middle] < target { // 中间的元素比目标小时，查找右半边
			low = middle + 1
		} else { // 处理等于target的情况
			if middle == 0 {
				return 0
			}
			// middel左边不等于target时，middle即为最左位
			if nums[middle-1] != target {
				return middle
			} else {
				// middel左边等于target时，high左移到middle-1位
				high = middle - 1
			}
		}
	}
	// 没有找到，返回-1
	return -1
}

func findRightIndex(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		middle := (low + high) / 2
		if nums[middle] > target { // 中间的元素比目标大时，查找左半边
			high = middle - 1
		} else if nums[middle] < target { // 中间的元素比目标小时，查找右半边
			low = middle + 1
		} else { // 处理等于target的情况
			if middle == len(nums)-1 {
				return len(nums) - 1
			}
			// middel右边不等于target时，middle即为最右位
			if nums[middle+1] != target {
				return middle
			} else {
				// middel右边等于target时，low右移到middle+1位
				low = middle + 1
			}
		}
	}
	return -1
}
