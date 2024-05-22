package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const (
	contextLines      = 1    // Showing lines before and after matches
	sampleSize        = 1024 // Number of bytes to read for content detection
	maxPrintableChars = 0.8  // Percentage threshold for printable characters to consider a file as a text file
)

func main() {
	searchStr := flag.String("s", "", "String to search for")
	searchDir := flag.String("d", ".", "Directory to search in")
	flag.Parse()

	if *searchStr == "" {
		fmt.Println("Please specify a string to search for")
		os.Exit(1)
	}

	err := filepath.Walk(*searchDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && isTextFile(path) {
			searchFile(path, *searchStr)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error scanning directory: %v\n", err)
		os.Exit(1)
	}
}
