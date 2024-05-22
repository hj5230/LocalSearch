package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func contains(s, substr string) bool {
	return bytes.Contains([]byte(s), []byte(substr))
}

func printContext(filePath string, matchedLineNum int, matchedLines map[int]string, searchStr string) {
	startLine := matchedLineNum - contextLines
	if startLine < 1 {
		startLine = 1
	}
	endLine := matchedLineNum + contextLines
	for i := startLine; i <= endLine; i++ {
		if i == matchedLineNum {
			fmt.Printf("%d | %s\n", i, highlightMatch(matchedLines[i], searchStr, i))
		} else {
			fmt.Printf("%d | %s\n", i, getLineContent(filePath, i))
		}
	}
}

func getLineContent(filePath string, lineNum int) string {
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentLine := 0
	for scanner.Scan() {
		currentLine++
		if currentLine == lineNum {
			return scanner.Text()
		}
	}
	return ""
}

func highlightMatch(line, searchStr string, lineNum int) string {
	index := bytes.Index([]byte(line), []byte(searchStr))
	if index != -1 {
		// Calculate the number of characters before the actual line content
		prefixLength := len(fmt.Sprintf("%d | ", lineNum))
		highlight := strings.Repeat("^", utf8.RuneCountInString(searchStr))
		return fmt.Sprintf("%s\n%s%s", line, strings.Repeat(" ", prefixLength+index), highlight)
	}
	return line
}
