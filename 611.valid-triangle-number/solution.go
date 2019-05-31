package leetcode

// brute force
// time complexity: O(n^3)
func triangleNumber(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if checkTriangle(nums[i], nums[j], nums[k]) {
					count++
				}
			}
		}
	}
	return count
}

func checkTriangle(a int, b int, c int) bool {
	return a+b > c && a+c > b && b+c > a
}
