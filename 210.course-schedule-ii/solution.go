package leetcode

func findOrder(numCourses int, prerequisites [][]int) []int {
	res := make([]int, 0)
	q := make([]int, 0)

	indegreeMap := make(map[int]int, 0)
	preMap := make(map[int][]int, 0)
	reachMap := make(map[int]bool, 0)
	for _, v := range prerequisites {
		indegreeMap[v[1]]++
		if preMap[v[0]] == nil {
			preMap[v[0]] = make([]int, 0)
		}
		preMap[v[0]] = append(preMap[v[0]], v[1])
	}
	// initialize the queue with degree zero
	for i := 0; i < numCourses; i++ {
		_, exist := indegreeMap[i]
		if !exist {
			q = append(q, i)
			reachMap[i] = true
		}
	}
	// then loop until all course numbers filled
	for len(q) > 0 {
		n := q[0]
		res = append(res, n)
		q = q[1:]
		for _, v := range preMap[n] {
			if reachMap[v] {
				// skip the reached number
				continue
			}
			_, exist := indegreeMap[v]
			indegreeMap[v]--
			if !exist || indegreeMap[v] == 0 {
				q = append(q, v)
				reachMap[v] = true
			}
		}
	}
	// return emtpy if impossible
	if len(res) < numCourses {
		return []int{}
	}
	// reverse the result
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
