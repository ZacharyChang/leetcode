package leetcode

func twoSum(numbers []int, target int) []int {
	for i := 0; i+1 < len(numbers); i++ {
		targetPos := binarySearch(numbers, i+1, target-numbers[i])
		if targetPos > 0 {
			return []int{
				i + 1, targetPos + 1,
			}
		}
	}
	return []int{0, 0}
}

func binarySearch(numbers []int, from int, target int) int {
	if len(numbers) == 0 || target < numbers[0] || target > numbers[len(numbers)-1] {
		return -1
	}
	start, end := from, len(numbers)-1
	if numbers[start] == target {
		return start
	}
	if numbers[end] == target {
		return end
	}
	for start < end-1 {
		mid := (start + end) / 2
		if numbers[mid] == target {
			return mid
		} else if numbers[mid] > target {
			end = mid
		} else {
			start = mid
		}
	}
	return -1
}
