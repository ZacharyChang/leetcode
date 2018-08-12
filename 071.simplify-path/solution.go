package leetcode

import (
	"strings"
)

func simplifyPath(path string) string {
	pathArray := strings.Split(path, "/")
	for i := 0; i < len(pathArray); {
		if pathArray[i] == "." || pathArray[i] == "" {
			pathArray = append(pathArray[:i], pathArray[i+1:]...)
		} else if pathArray[i] == ".." {
			if i-1 >= 0 {
				pathArray = append(pathArray[:i-1], pathArray[i+1:]...)
				i--
			} else {
				pathArray = pathArray[i+1:]
			}
		} else {
			i++
		}
	}
	return "/" + strings.Join(pathArray, "/")
}
