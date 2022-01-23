// Leet #937
// Reorder Data in Logs
// start: Friday, October 29, 2021
// finish: Friday, November 5, 2021

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//logs := []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}
	logs := []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo", "a2 act car"}
	sortedLogs := reorderLogFiles(logs)
	fmt.Println("sortedLogs =", sortedLogs)
}

func reorderLogFiles(logs []string) []string {
	// slices
	digitLogs := []string{}
	letterLogs := []string{}

	// separate digit logs from letter logs
	for _, log := range logs {
		// logType represents the first character (letter or number) in the
		// second "word" of each log
		logType := strings.Fields(log)[1][0]

		if logType >= '0' && logType <= '9' {
			digitLogs = append(digitLogs, log)
		} else {
			letterLogs = append(letterLogs, log)
		}
	}

	// SliceStable
	// 1. lexiographically by content
	// 2. by identifier - implicitly by identifier
	sort.SliceStable(letterLogs, func(i, j int) bool {
		// first empty space in log represents start of log message
		iIndex := strings.Index(letterLogs[i], " ")
		jIndex := strings.Index(letterLogs[j], " ")

		iLog := letterLogs[i][iIndex+1:]
		jLog := letterLogs[j][jIndex+1:]

		// if logs equal, sort by identifier
		if iLog == jLog {
			return letterLogs[i][0] < letterLogs[j][0]
		}
		// else, sort alphabetically by log message
		return iLog < jLog
	})

	return append(letterLogs, digitLogs...)
}
