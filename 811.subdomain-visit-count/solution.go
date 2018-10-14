package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	domainMap := make(map[string]int, len(cpdomains))
	result := make([]string, 0)
	for _, v := range cpdomains {
		arr := strings.Split(v, " ")
		count, _ := strconv.Atoi(arr[0])
		domain := arr[1]

		domainMap[domain] += count
		for strings.IndexByte(domain, '.') >= 0 {
			index := strings.IndexByte(domain, '.')
			domainMap[domain[index+1:]] += count
			domain = domain[index+1:]
		}
	}
	for k, v := range domainMap {
		result = append(result, fmt.Sprintf("%d %s", v, k))
	}
	return result
}
