package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if nums1[j] > nums2[i] {
				copy(nums1[j+1:], nums1[j:])
				nums1[j] = nums2[i]
				m++
				break
			}
			// 若nums1末尾元素小于nums2后元素，复制nums2到nums1末尾
			if j == m-1 {
				copy(nums1[j+1:], nums2[i:])
				return
			}
		}
	}
}
