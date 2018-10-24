package leetcode

func findClosestElements(arr []int, k int, x int) []int {
	start := 0
	end := len(arr) - k
	for start < end {
		mid := (start + end) / 2
		if x-arr[mid] > arr[mid+k]-x {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return arr[start : start+k]
}
