package leetcode

func canFinish(numCourses int, prerequisites [][]int) bool {

	queue := make([]int, 0)
	indegreeMap := make(map[int]int, 0)
	preMap := make(map[int][]int, 0)
	reachMap := make(map[int]bool, 0)

	for _, v := range prerequisites {
		indegreeMap[v[1]]++
		preMap[v[0]] = append(preMap[v[0]], v[1])
	}

	for i := 0; i < numCourses; i++ {
		if indegreeMap[i] <= 0 {
			queue = append(queue, i)
			reachMap[i] = true
		}
	}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		for _, v := range preMap[n] {
			indegreeMap[v]--
			if indegreeMap[v] == 0 && !reachMap[v] {
				queue = append(queue, v)
				reachMap[v] = true
			}
		}
	}
	return len(reachMap) == numCourses
}
