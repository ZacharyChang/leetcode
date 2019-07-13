package leetcode

// brute force
// todo: stack
func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		for j := i + 1; j < len(T); j++ {
			if T[j] > T[i] && res[i] == 0 {
				res[i] = j - i
				break
			}
		}
	}
	return res
}
