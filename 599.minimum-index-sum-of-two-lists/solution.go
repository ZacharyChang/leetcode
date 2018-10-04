package leetcode

func findRestaurant(list1 []string, list2 []string) []string {
	result := make([]string, 0)
	minIndexSum := len(list1) + len(list2) + 1
	restMap := make(map[string]int, 0)
	for index, v := range list1 {
		restMap[v] = index + 1
	}
	for index, v := range list2 {
		if restMap[v] > 0 {
			if restMap[v]+index < minIndexSum {
				result = []string{
					v,
				}
				minIndexSum = restMap[v] + index
			} else if restMap[v]+index == minIndexSum {
				result = append(result, v)
			}
		}
	}
	return result
}
