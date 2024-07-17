package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func SearchFile(filePath, searchStr string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	if !isLikelyText(content) {
		return
	}

	scanner := bufio.NewScanner(bytes.NewReader(content))
	lineNum := 0
	matchedLines := make(map[int]string)

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if contains(line, searchStr) {
			matchedLines[lineNum] = line
		}
	}

	if len(matchedLines) > 0 {
		fmt.Printf("- %s\n", filepath.Base(filePath))
		if contains(filepath.Base(filePath), searchStr) {
			fmt.Println("---")
		}
		fmt.Printf("* %s:\n", filePath)
		for lineNum := range matchedLines {
			printContext(filePath, lineNum, matchedLines, searchStr)
		}
		fmt.Println("---")
	}
}
