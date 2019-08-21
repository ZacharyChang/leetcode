package leetcode

import (
	"strings"
)

func numUniqueEmails(emails []string) int {
	emailMap := make(map[string]int, 0)
	for _, v := range emails {
		emailIdx := strings.Index(v, "@")
		localName := v[:emailIdx]
		domainName := v[emailIdx+1:]
		localName = strings.Replace(localName, ".", "", -1)
		if strings.Contains(localName, "+") {
			localName = localName[:strings.Index(localName, "+")]
		}
		emailMap[localName+"@"+domainName]++
	}
	return len(emailMap)
}
