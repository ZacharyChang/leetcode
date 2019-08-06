package leetcode

func eventualSafeNodes(graph [][]int) []int {
	isSafe := make(map[int]bool, 0)

	for {
		hasUpdated := false
		for i := 0; i < len(graph); i++ {
			if isSafe[i] {
				continue
			}
			// check graph[i] is safe or not
			// by current safe nodes
			safe := true
			for _, v := range graph[i] {
				if !isSafe[v] {
					safe = false
					break
				}
			}
			if safe {
				isSafe[i] = true
				hasUpdated = true
			}
		}
		// return result when all nodes are update to date
		// else continue looping again
		if !hasUpdated {
			res := make([]int, 0)
			for i := 0; i < len(graph); i++ {
				if isSafe[i] {
					res = append(res, i)
				}
			}
			return res
		}
	}
}
