package leetcode

// merge array and get median
// time complexity: O(m+n)
// TODO: time complexity to O(log(m+n))
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := make([]int, 0)
	m, n := len(nums1), len(nums2)
	mid := (m + n) / 2
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
		if len(merged) > mid {
			goto RET
		}
	}
	if i == m {
		merged = append(merged, nums2[j:]...)
	}
	if j == n {
		merged = append(merged, nums1[i:]...)
	}
RET:
	if (m+n)%2 == 1 {
		return float64(merged[mid])
	}
	return float64(merged[mid-1]+merged[mid]) / 2
}
