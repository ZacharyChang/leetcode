package leetcode

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	letterLogs := make(letterLogList, 0)
	digitLogs := make([]string, 0)
	for _, log := range logs {
		word := strings.Split(log, " ")[1]
		if isNumber(word) {
			digitLogs = append(digitLogs, log)
		} else {
			letterLogs = append(letterLogs, log)
		}
	}
	sort.Sort(letterLogs)
	return append(letterLogs, digitLogs...)
}

func isNumber(str string) bool {
	return str[0] >= '0' && str[0] <= '9'
}

type letterLogList []string

func (l letterLogList) Len() int {
	return len(l)
}

func (l letterLogList) Less(i, j int) bool {
	idxi := strings.Index(l[i], " ")
	idxj := strings.Index(l[j], " ")
	li := l[i][idxi:]
	lj := l[j][idxj:]
	if li != lj {
		return li < lj
	}
	return l[i][:idxi] < l[j][:idxj]
}

func (l letterLogList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
