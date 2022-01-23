// Leet #937
// Reorder Data in Logs
// Friday, October 29, 2021

package main

import (
	"fmt"
	"sort"
	"strings"
)

// Tried to create and use a custom LetterLog type
type LetterLog struct {
	Identifier string
	Message    []string
}

type ByMessage []LetterLog

func (a ByMessage) Len() int {
	return len(a)
}

func (a ByMessage) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByMessage) Less(i, j int) bool {
	return a[i].Message < a[j].Message
}


func main() {
	//logs := []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}
	logs := []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo", "a2 act car"}
	sortedLogs := reorderLogFiles(logs)
	fmt.Println("sortedLogs =", sortedLogs)
}

// sorts an array of log messages based on provided rules
func reorderLogFiles(logs []string) []string {
	digitLogs := []string{}
	letterLogs := []LetterLog{}

	for _, log := range logs {
		// logType represents the first character (letter or number)
		// in the second "word" of each log
		logType := strings.Fields(log)[1][0]

		if logType >= '0' && logType <= '9' {
			digitLogs = append(digitLogs, log)
		} else {
			identifier := strings.Fields(log)[0]
			message := strings.Fields(log)[1:]
			letterLogs = append(letterLogs, LetterLog{identifier, message})
		}
	}

	fmt.Println("letterLogs before =", letterLogs)

	sort.SliceStable(letterLogs, func(i, j int) bool {
		iString := strings.Join(letterLogs[i].Message, " ")
		jString := strings.Join(letterLogs[j].Message, " ")

		fmt.Println("iString =", iString)
		fmt.Println("jString =", jString)

		if iString == jString {
			return letterLogs[i].Identifier < letterLogs[j].Identifier
		}
		return iString < jString
	})

	for _, log := range letterLogs {
		log := log.Identifier + strings.Fields(log.Message[])
	}


	fmt.Println("letterLogs after =", letterLogs)
	return digitLogs
}
