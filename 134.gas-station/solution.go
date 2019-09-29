package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	diff := make([]int, len(gas))
	for i := range gas {
		diff[i] = gas[i] - cost[i]
	}
	for i := 0; i < len(gas); i++ {
		if check(diff, i) {
			return i
		}
	}
	return -1
}

func check(diff []int, index int) bool {
	sum := diff[index]
	if sum < 0 {
		return false
	}
	for i := index + 1; i%len(diff) != index; i++ {
		sum += diff[i%len(diff)]
		if sum < 0 {
			return false
		}
	}
	return true
}
