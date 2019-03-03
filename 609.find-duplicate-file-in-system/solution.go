package leetcode

import "strings"

func findDuplicate(paths []string) [][]string {
	contentMap := make(map[string][]string, 0)
	res := make([][]string, 0)
	for _, path := range paths {
		splitArray := strings.Split(path, " ")
		dir := splitArray[0]
		for i := 1; i < len(splitArray); i++ {
			index := strings.Index(splitArray[i], "(")
			filename := splitArray[i][0:index]
			content := splitArray[i][index+1 : len(splitArray[i])]
			contentMap[content] = append(contentMap[content], dir+"/"+filename)
		}
	}
	for _, v := range contentMap {
		if len(v) > 1 {
			res = append(res, v)
		}
	}
	return res
}
